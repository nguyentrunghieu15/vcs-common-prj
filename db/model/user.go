package model

import (
	"gorm.io/gorm"
)

type UserRole int

const (
	RoleAdmin UserRole = iota
	RoleUser
)

var role = []string{"admin", "user"}

func (c *UserRole) String() string {
	return role[*c]
}

type BaseModel struct {
	gorm.Model
	CreatedBy int
	UpdatedBy int
	DeletedBy int
}

type User struct {
	BaseModel
	Email         string
	FullName      string
	Phone         string
	Avatar        string
	IsSupperAdmin bool
	Roles         string
}
