package share

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TextRequest 定义接口请求参数结构体
type TextRequest struct {
	Content     string `json:"content" query:"content" validate:"required,min=1,max=1000"`
	Title       string `json:"title" query:"title" validate:"omitempty,min=1,max=100"`
	Type        string `json:"type" query:"type" validate:"omitempty,oneof=plain html markdown"`
	ExpireValue string `json:"expire_value" query:"expire_value" validate:"omitempty,gt=0"`
	ExpireType  string `json:"expire_type" query:"expire_type" validate:"omitempty,oneof=minute hour day week month year forever count"`
}

// TextResponse 定义接口响应参数结构体
type TextResponse struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	ExpireValue string `json:"expire_value"`
	ExpireType  string `json:"expire_type"`
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

	// 设置 ExpireValue 的缺省值
	// 不能直接tag default
	// https://github.com/go-playground/validator/issues/263
	if req.ExpireValue == "" {
		req.ExpireValue = "1"
	}

	if req.Type == "" {
		req.Type = "plain"
	}

	if req.ExpireType == "" {
		req.ExpireType = "day"
	}

	// 构造响应
	resp := &TextResponse{
		ID:          "text-123", // TODO: 生成唯一ID
		Content:     req.Content,
		Title:       req.Title,
		Type:        req.Type,
		ExpireValue: req.ExpireValue,
		ExpireType:  req.ExpireType,
	}

	return c.JSON(http.StatusOK, resp)
}
