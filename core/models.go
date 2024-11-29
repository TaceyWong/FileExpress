package core

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
	Code         string     `gorm:"uniqueIndex;max_length:255"`       // 分享码
	Prefix       string     `gorm:"max_length:255;default:''"`        // 前缀
	Suffix       string     `gorm:"max_length:255;default:''"`        // 后缀
	UUIDFileName string     `gorm:"max_length:255;default:''"`        // uuid文件名
	FilePath     string     `gorm:"max_length:255;default:'';isnull"` // 文件路径
	FileSize     int64      `gorm:"default:0"`                        // 文件大小
	Title        string     `gorm:"max_length:255;default:''"`        // 分享标题
	Text         string     `gorm:"isnull"`                           // 文本内容
	ExpiredAt    *time.Time `gorm:"isnull"`                           // 过期时间
	ExpiredCount int        `gorm:"default:0"`                        // 过期次数
	UsedCount    int        `gorm:"default:0"`                        // 使用次数
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
	Key   string `gorm:"uniqueIndex;max_length:255"` // 键
	Value string `gorm:"isnull"`                     // 值
}
