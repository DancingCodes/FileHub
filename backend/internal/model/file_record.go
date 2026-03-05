package model

import "time"

type FileRecord struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FileName  string    `gorm:"type:varchar(255)" json:"file_name"`
	FileUUID  string    `gorm:"type:varchar(100);uniqueIndex" json:"file_uuid"`
	FileSize  int64     `json:"file_size"`
	FileType  string    `gorm:"type:varchar(50)" json:"file_type"`
	FileUrl   string    `gorm:"type:varchar(100);uniqueIndex" json:"file_url"`
	CreatedAt time.Time `json:"created_at"`
}
