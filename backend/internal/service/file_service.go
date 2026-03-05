package service

import (
	"backend/internal/db"
	"backend/internal/model"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func SaveFileRecord(file *multipart.FileHeader) (*model.FileRecord, error) {
	ext := filepath.Ext(file.Filename)
	uniqueName := uuid.New().String() + ext

	baseURL := os.Getenv("IMG_BASE_URL")

	record := model.FileRecord{
		FileName: baseURL + file.Filename,
		FileUUID: uniqueName,
		FileSize: file.Size,
		FileType: ext,
	}

	if err := db.DB.Create(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}
