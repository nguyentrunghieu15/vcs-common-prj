package model

import (
	"time"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindOneById(int) (*User, error)
	FindOneByEmail(string) (*User, error)
	CreateUser(map[string]interface{}) (*User, error)
	UpdateOneByEmail(string, map[string]interface{}) (*User, error)
	UpdateOneById(int, map[string]interface{}) (*User, error)
	DeleteOneById(int) error
	DeleteOneByEmail(string) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (c *UserRepository) FindOneById(id int) (*User, error) {
	var user = User{}
	if err := c.db.First(&user, id); err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (c *UserRepository) FindOneByEmail(email string) (*User, error) {
	var user = User{}
	if err := c.db.First(&user, "email = ?", email); err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (c *UserRepository) CreateUser(u map[string]interface{}) (*User, error) {
	var user User
	u["created_at"] = time.Now()
	result := c.db.Model(&user).Create(u)
	if result.Error != nil {
		return &user, result.Error
	}
	c.db.Where("email = ?", u["email"]).First(&user)
	return &user, result.Error
}

func (c *UserRepository) UpdateOneByEmail(email string, u map[string]interface{}) (*User, error) {
	var user User
	u["updated_at"] = time.Now()
	result := c.db.Model(&user).Where("email = ?", email).Updates(u)
	if result.Error != nil {
		return &user, result.Error
	}
	c.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (c *UserRepository) UpdateOneById(id int, u map[string]interface{}) (*User, error) {
	var user User
	u["updated_at"] = time.Now()
	result := c.db.Model(&user).Where("id = ?", id).Updates(u)
	if result.Error != nil {
		return &user, result.Error
	}
	c.db.First(&user, id)
	return &user, result.Error
}

func (c *UserRepository) DeleteOneById(id int) error {
	result := c.db.Where("id = ?", id).Delete(&User{})
	return result.Error
}
func (c *UserRepository) DeleteOneByEmail(email string) error {
	result := c.db.Where("email = ?", email).Delete(&User{})
	return result.Error
}
