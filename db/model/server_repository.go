package model

import (
	"time"

	"gorm.io/gorm"
)

type ServerRepositoryDecorator interface {
	FindOneById(int) (*Server, error)
	FindOneByName(string) (*Server, error)
	CreateServer(map[string]interface{}) (*Server, error)
	UpdateOneByName(string, map[string]interface{}) (*Server, error)
	UpdateOneById(int, map[string]interface{}) (*Server, error)
	DeleteOneById(int) error
	DeleteOneByName(string) error
}

type ServerRepository struct {
	db *gorm.DB
}

func CreateServerRepository(db *gorm.DB) *ServerRepository {
	return &ServerRepository{db: db}
}

func (c *ServerRepository) FindOneById(id int) (*Server, error) {
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
	s["created_at "] = time.Now()
	result := c.db.Model(&server).Create(s)
	return &server, result.Error
}

func (c *ServerRepository) UpdateOneByName(name string, s map[string]interface{}) (*Server, error) {
	var server Server
	s["updated_at "] = time.Now()
	result := c.db.Model(&server).Where("name = ?", name).Updates(s)
	return &server, result.Error
}

func (c *ServerRepository) UpdateOneById(id int, s map[string]interface{}) (*Server, error) {
	var server Server
	s["updated_at "] = time.Now()
	result := c.db.Model(&server).Where("id = ?", id).Updates(s)
	return &server, result.Error
}
func (c *ServerRepository) DeleteOneById(id int) error {
	result := c.db.Where("id = ?", id).Delete(&Server{})
	return result.Error
}
func (c *ServerRepository) DeleteOneByName(name string) error {
	result := c.db.Where("name = ?", name).Delete(&Server{})
	return result.Error
}
