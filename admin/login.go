package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

// Dashboard 处理管理员仪表盘请求
func Dashboard(c echo.Context) error {
	return c.String(http.StatusOK, "dashboard")
}

// DeleteFile 处理删除文件请求
func DeleteFile(c echo.Context) error {
	return c.String(http.StatusOK, "delete file")
}

// FileList 处理文件列表请求
func FileList(c echo.Context) error {
	return c.String(http.StatusOK, "file list")
}

// GetConfig 处理获取配置请求
func GetConfig(c echo.Context) error {
	return c.String(http.StatusOK, "get config")
}

// UpdateConfig 处理更新配置请求
func UpdateConfig(c echo.Context) error {
	return c.String(http.StatusOK, "update config")
}

// FileDownload 处理文件下载请求
func FileDownload(c echo.Context) error {
	return c.String(http.StatusOK, "file download")
}

// GetLocalList 处理本地文件列表请求
func GetLocalList(c echo.Context) error {
	return c.String(http.StatusOK, "local list")
}

// DeleteLocalFile 处理删除本地文件请求
func DeleteLocalFile(c echo.Context) error {
	return c.String(http.StatusOK, "delete local file")
}

// ShareLocalFile 处理分享本地文件请求
func ShareLocalFile(c echo.Context) error {
	return c.String(http.StatusOK, "share local file")
}
