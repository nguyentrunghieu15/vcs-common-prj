package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	CreatedBy int            `json:"createdBy"`
	UpdatedBy int            `json:"updatedBy"`
	DeletedBy int            `json:"deletedBy"`
}

type User struct {
	BaseModel
	Email         string   `json:"email"`
	FullName      string   `json:"fullName"`
	Phone         string   `json:"phone"`
	Avatar        string   `json:"avatar"`
	IsSupperAdmin bool     `json:"isSupperAdmin"`
	Roles         UserRole `json:"roles"`
	Password      string   `json:"password"`
}
