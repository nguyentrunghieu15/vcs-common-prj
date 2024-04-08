package managedb

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionContext interface {
	Connect(string) (interface{}, error)
}

type PostgreContext struct{}

func (c *PostgreContext) Connect(dsn string) (interface{}, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

type RedisContext struct{}

func (c *RedisContext) Connect(dsn string) (interface{}, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}

type ElasticContext struct{}

type ElasticConfigDsn struct {
	Addresses  []string
	Username   string
	Password   string
	PathCACert string
}

func (c *ElasticContext) Connect(dsn string) (interface{}, error) {
	var config ElasticConfigDsn
	err := json.Unmarshal([]byte(dsn), &config)
	if err != nil {
		return nil, fmt.Errorf("Undefine dsn elastic: %v", err)
	}

	cert, _ := os.ReadFile(config.PathCACert)

	cfg := elasticsearch.Config{
		Addresses: config.Addresses,
		Username:  config.Username,
		Password:  config.Password,
		CACert:    cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	return es, err
}

type Connection struct {
	Context ConnectionContext
	Dsn     string
}

type ManagedbConnection struct {
	mu          sync.Mutex
	connections map[string]interface{}
}

var m ManagedbConnection = ManagedbConnection{
	connections: map[string]interface{}{},
}

func GetConnection(con Connection) (interface{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c, ok := m.connections[con.Dsn]
	if ok {
		return c, nil
	}
	connection, err := con.Context.Connect(con.Dsn)
	if err != nil {
		return connection, err
	}
	m.connections[con.Dsn] = connection
	return connection, err
}
