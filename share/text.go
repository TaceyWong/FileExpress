package share

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TextRequest 定义接口请求参数结构体
type TextRequest struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
	Title   string `json:"title" validate:"required,min=1,max=100"`
	Type    string `json:"type" validate:"required,oneof=plain html markdown"`
}

// TextResponse 定义接口响应参数结构体
type TextResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

// Text 处理文本分享请求
func Text(c echo.Context) error {
	// 绑定并验证请求参数
	req := new(TextRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 构造响应
	resp := &TextResponse{
		ID:      "text-123", // TODO: 生成唯一ID
		Content: req.Content,
		Title:   req.Title,
		Type:    req.Type,
	}

	return c.JSON(http.StatusOK, resp)
}
