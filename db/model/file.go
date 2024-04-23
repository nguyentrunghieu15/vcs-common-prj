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
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	CreatedBy int            `json:"createdBy"`
	UpdatedBy int            `json:"updatedBy"`
	DeletedBy int            `json:"deletedBy"`
	FileName  string         `json:"fileName"`
	FilePath  string         `json:"filePath"`
	FileSize  int            `json:"fileSize"`
	Owner     int            `json:"owner"`
	Status    FileStatus     `json:"status"`
}
