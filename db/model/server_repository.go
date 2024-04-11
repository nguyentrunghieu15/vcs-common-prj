package model

import (
	"gorm.io/gorm"
)

type ServerRepositoryDecorator interface {
	FindOneById(int) (*Server, error)
	FindOneByName(string) (*Server, error)
	CreateServer(map[string]interface{}) error
	UpdateOneByName(string, map[string]interface{}) error
	UpdateOneById(int, map[string]interface{}) error
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

func (c *ServerRepository) CreateServer(s map[string]interface{}) error {
	result := c.db.Model(&Server{}).Create(s)
	return result.Error
}

func (c *ServerRepository) UpdateOneByName(name string, s map[string]interface{}) error {
	result := c.db.Model(&Server{}).Where("name = ?", name).Updates(s)
	return result.Error
}

func (c *ServerRepository) UpdateOneById(id int, s map[string]interface{}) error {
	result := c.db.Model(&Server{}).Where("id = ?", id).Updates(s)
	return result.Error
}
func (c *ServerRepository) DeleteOneById(id int) error {
	result := c.db.Where("id = ?", id).Delete(&Server{})
	return result.Error
}
func (c *ServerRepository) DeleteOneByName(name string) error {
	result := c.db.Where("name = ?", name).Delete(&Server{})
	return result.Error
}
