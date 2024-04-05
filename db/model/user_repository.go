package model

import (
	"gorm.io/gorm"
)

type UserRepositoryDecorator interface {
	FindOneById(int) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (c *UserRepository) FindOneById(id int) (*User, error) {
	var user = User{}
	if err := c.db.First(&user, id); err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}
