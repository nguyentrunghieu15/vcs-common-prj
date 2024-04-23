package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IFileRepository interface {
	CreateFile(map[string]interface{}) error
	DeleteFileById(uuid.UUID) error
	FindAllFileOfUser(int) (*[]File, error)
	FindFileById(uuid.UUID) (*File, error)
	UpdateFileById(uuid.UUID, map[string]interface{}) (*File, error)
}

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (f *FileRepository) FindFileById(id uuid.UUID) (*File, error) {
	var file File
	result := f.db.First(&file, id)
	return &file, result.Error
}

func (f *FileRepository) UpdateFileById(id uuid.UUID, data map[string]interface{}) (*File, error) {
	result := f.db.Model(&File{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, result.Error
	}
	var file File
	f.db.First(&file, id)
	return &file, nil
}

func (f *FileRepository) CreateFile(data map[string]interface{}) error {
	result := f.db.Model(&File{}).Create(data)
	return result.Error
}

func (f *FileRepository) DeleteFileById(id uuid.UUID) error {
	result := f.db.Delete(&File{}, id)
	return result.Error
}

func (f *FileRepository) FindAllFileOfUser(userId int) (*[]File, error) {
	var files []File
	result := f.db.Where("owner = ?", userId).Find(&files)
	return &files, result.Error
}
