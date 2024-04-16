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
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedBy int
	UpdatedBy int
	DeletedBy int
	Name      string
	Status    ServerStatus
	Ipv4      string
}
