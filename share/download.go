package share

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// DownloadRequest 定义接口请求参数结构体
type DownloadRequest struct {
	ID string `json:"id" validate:"required"`
}

// DownloadResponse 定义接口响应参数结构体
type DownloadResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

// Download 处理文本分享请求
func Download(c echo.Context) error {
	// 绑定并验证请求参数
	req := new(DownloadRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 构造响应
	resp := &DownloadResponse{
		ID: "download-123", // TODO: 生成唯一ID
	}

	return c.JSON(http.StatusOK, resp)
}
