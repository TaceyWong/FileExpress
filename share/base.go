package share

import (
	"math/rand"
	"time"
)

// ExpireType 定义过期时间类型枚举
type ExpireType string

const (
	ExpireTypeMinute  ExpireType = "minute"
	ExpireTypeHour    ExpireType = "hour"
	ExpireTypeDay     ExpireType = "day"
	ExpireTypeWeek    ExpireType = "week"
	ExpireTypeMonth   ExpireType = "month"
	ExpireTypeYear    ExpireType = "year"
	ExpireTypeForever ExpireType = "forever"
	ExpireTypeCount   ExpireType = "count"
)

type ExpireInfo struct {
	ExpireAt    time.Time
	ExpireCount int
	UsedCount   int
	Code        string
}

func init() {
	// rand.Seed(time.Now().UnixNano())
	// Go 1.20+ 的 math/rand 包会自动初始化随机数生成器，不需要手动设置种子。

}

func GetExpireInfo(expireValue int, expireType ExpireType) *ExpireInfo {
	now := time.Now()
	expireInfo := &ExpireInfo{
		ExpireCount: -1,
		UsedCount:   0,
		Code:        "",
		ExpireAt:    now,
	}
	expireInfo.Code = GetRandomCode(6)
	switch expireType {
	case ExpireTypeMinute:
		expireInfo.ExpireAt = expireInfo.ExpireAt.Add(time.Duration(expireValue) * time.Minute)
	case ExpireTypeHour:
		expireInfo.ExpireAt = expireInfo.ExpireAt.Add(time.Duration(expireValue) * time.Hour)
	case ExpireTypeDay:
		expireInfo.ExpireAt = expireInfo.ExpireAt.Add(time.Duration(expireValue) * time.Hour * 24)
	case ExpireTypeWeek:
		expireInfo.ExpireAt = expireInfo.ExpireAt.AddDate(0, 0, expireValue)
	case ExpireTypeMonth:
		expireInfo.ExpireAt = expireInfo.ExpireAt.AddDate(0, expireValue, 0)
	case ExpireTypeYear:
		expireInfo.ExpireAt = expireInfo.ExpireAt.AddDate(expireValue, 0, 0)
	case ExpireTypeForever:
		expireInfo.Code = GetRandomCode(8)
	default:
		expireInfo.ExpireAt = expireInfo.ExpireAt.AddDate(0, 0, 1)
	}

	return expireInfo
}

// GetRandomCode 生成随机码
func GetRandomCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
