package model

import (
	"gorm.io/gorm"
)

type UserRepositoryDecorator interface {
	FindOneById(int) (*User, error)
	FindOneByEmail(string) (*User, error)
	CreateUser(*User) error
	UpdateOneByEmail(string, *User) error
	UpdateOneById(int, *User) error
	DeleteOneById(int) error
	DeleteOneByEmail(string) error
}

type UserRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
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

func (c *UserRepository) CreateUser(u *User) error {
	result := c.db.Create(u)
	return result.Error
}

func (c *UserRepository) UpdateOneByEmail(email string, u *User) error {
	result := c.db.Model(&User{}).Where("email = ?", email).Updates(map[string]interface{}{
		"full_name": u.FullName,
		"phone":     u.Phone,
		"avatar":    u.Avatar,
		"roles":     u.Roles,
	})
	return result.Error
}

func (c *UserRepository) UpdateOneById(id int, u *User) error {
	result := c.db.Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"full_name": u.FullName,
		"phone":     u.Phone,
		"avatar":    u.Avatar,
		"roles":     u.Roles,
	})
	return result.Error
}
func (c *UserRepository) DeleteOneById(id int) error {
	result := c.db.Where("id = ?", id).Delete(&User{})
	return result.Error
}
func (c *UserRepository) DeleteOneByEmail(email string) error {
	result := c.db.Where("email = ?", email).Delete(&User{})
	return result.Error
}
