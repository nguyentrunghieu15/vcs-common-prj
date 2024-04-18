package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileStatus string

const (
	Exported  FileStatus = "exported"
	Exporting FileStatus = "exporting"
	Fail      FileStatus = "fail"
)

type File struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedBy int
	UpdatedBy int
	DeletedBy int
	FileName  string
	FilePath  string
	FileSize  int
	Owner     int
	Status    FileStatus
}
