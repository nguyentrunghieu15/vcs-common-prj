package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServerRepositoryDecorator interface {
	FindOneById(uuid.UUID) (*Server, error)
	FindOneByName(string) (*Server, error)
	CreateServer(map[string]interface{}) (*Server, error)
	UpdateOneByName(string, map[string]interface{}) (*Server, error)
	UpdateOneById(uuid.UUID, map[string]interface{}) (*Server, error)
	DeleteOneById(uuid.UUID) error
	DeleteOneByName(string) error
}

type ServerRepository struct {
	db *gorm.DB
}

func CreateServerRepository(db *gorm.DB) *ServerRepository {
	return &ServerRepository{db: db}
}

func (c *ServerRepository) FindOneById(id uuid.UUID) (*Server, error) {
	var Server = Server{}
	if err := c.db.First(&Server, id); err.Error != nil {
		return nil, err.Error
	}
	return &Server, nil
}

func (c *ServerRepository) FindOneByName(name string) (*Server, error) {
	var Server = Server{}
	if err := c.db.First(&Server, "name = ?", name); err.Error != nil {
		return nil, err.Error
	}
	return &Server, nil
}

func (c *ServerRepository) CreateServer(s map[string]interface{}) (*Server, error) {
	var server Server
	s["created_at"] = time.Now()
	result := c.db.Model(&server).Create(s)
	if result.Error != nil {
		return &server, result.Error
	}
	c.db.Where("name = ?", s["name"]).First(&server)
	return &server, result.Error
}

func (c *ServerRepository) UpdateOneByName(name string, s map[string]interface{}) (*Server, error) {
	var server Server
	s["updated_at"] = time.Now()
	result := c.db.Model(&server).Where("name = ?", name).Updates(s)
	if result.Error != nil {
		return &server, result.Error
	}
	c.db.Where("name = ?", name).First(&server)
	return &server, result.Error
}

func (c *ServerRepository) UpdateOneById(id uuid.UUID, s map[string]interface{}) (*Server, error) {
	var server Server
	s["updated_at"] = time.Now()
	result := c.db.Model(&Server{}).Where("id = ?", id).Updates(s)
	if result.Error != nil {
		return &server, result.Error
	}
	c.db.First(&server, id)
	return &server, result.Error
}
func (c *ServerRepository) DeleteOneById(id uuid.UUID) error {
	result := c.db.Where("id = ?", id).Delete(&Server{})
	return result.Error
}
func (c *ServerRepository) DeleteOneByName(name string) error {
	result := c.db.Where("name = ?", name).Delete(&Server{})
	return result.Error
}
