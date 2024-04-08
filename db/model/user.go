package model

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

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
	Roles         UserRole
	Password      string
}
