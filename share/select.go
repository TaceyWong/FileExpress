package share

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SelectRequest 定义接口请求参数结构体
type SelectRequest struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
	Title   string `json:"title" validate:"required,min=1,max=100"`
	Type    string `json:"type" validate:"required,oneof=plain html markdown"`
}

// SelectResponse 定义接口响应参数结构体
type SelectResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

// Select 处理选择分享请求
func Select(c echo.Context) error {
	// 绑定并验证请求参数
	req := new(SelectRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 构造响应
	resp := &SelectResponse{
		ID:      "select-123", // TODO: 生成唯一ID
		Content: req.Content,
		Title:   req.Title,
		Type:    req.Type,
	}

	return c.JSON(http.StatusOK, resp)
}

// GetCodeFileRequest 定义获取代码文件分享请求参数结构体
type GetCodeFileRequest struct {
	ID string `json:"id" validate:"required"`
}

// GetCodeFileResponse 定义获取代码文件分享响应参数结构体
type GetCodeFileResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

// GetCodeFile 处理获取代码文件分享请求
func GetCodeFile(c echo.Context) error {
	// 绑定并验证请求参数
	req := new(GetCodeFileRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 构造响应
	resp := &GetCodeFileResponse{
		ID:      req.ID,
		Content: "code file content",
		Title:   "code file title",
		Type:    "code",
	}
	return c.JSON(http.StatusOK, resp)
}
