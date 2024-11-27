package main

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type FileCodes struct {
	gorm.Model
	Code         int    // 分享码
	Prefix       string // 前缀
	Suffix       string // 后缀
	UUIDFileName string // 文件名
	FilePath     string // 文件路径
	FileSize     int64  // 文件大小
	Text         string
	ExpiredAt    *time.Time
	ExpiredCount int
	UsedCount    int
}

func (f *FileCodes) TableName() string {
	return "file_codes"
}

func (f *FileCodes) IsExpired() bool {
	// 按时间
	if f.ExpiredAt == nil {
		return false
	}
	if f.ExpiredAt != nil && f.ExpiredCount < 0 {
		return f.ExpiredAt.Before(time.Now())
	}
	// 按次数
	return f.ExpiredCount <= 0
}
func (f *FileCodes) GetFilePath() string {
	return f.FilePath + "/" + f.Prefix + "/" + f.UUIDFileName + f.Suffix
}

type KeyValue struct {
	gorm.Model
	Key   string
	Value string
}
