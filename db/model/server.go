package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServerStatus string

const (
	On  ServerStatus = "on"
	Off ServerStatus = "off"
)

type Server struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	CreatedBy int            `json:"createdBy"`
	UpdatedBy int            `json:"updatedBy"`
	DeletedBy int            `json:"deletedBy"`
	Name      string         `json:"name"`
	Status    ServerStatus   `json:"status"`
	Ipv4      string         `json:"ipv4"`
}
