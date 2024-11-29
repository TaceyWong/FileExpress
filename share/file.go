package share

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

// FileRequest 定义接口请求参数结构体
type FileRequest struct {
	ExpireValue int           `form:"expire_value" validate:"required,min=1,max=1000"`
	ExpireStyle string        `form:"expire_style" validate:"required,oneof=minute hour day week month year"`
	File        *echo.Context `form:"file" validate:"required"`
}

// File 处理文件分享请求
func File(c echo.Context) error {

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "文件上传失败")
	}
	maxSizeMb := 1024 * 1024 * 1024
	if file.Size > int64(maxSizeMb) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("大小超过限制,最大为%f MB",
			float64(maxSizeMb)/1024/1024))
	}

	// 获取过期时间和下载次数
	expireTime := c.FormValue("expire_time")
	expireCount, _ := strconv.Atoi(c.FormValue("expire_count"))

	// 保存文件
	dst := fmt.Sprintf("uploads/%s", file.Filename)
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "无法打开文件")
	}
	defer src.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "无法创建目标文件")
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "文件保存失败")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"name":         file.Filename,
		"expire_time":  expireTime,
		"expire_count": expireCount,
	})
}
