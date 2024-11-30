package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"

	"FileExpress/admin"
	"FileExpress/share"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 定义自定义claims结构
type JwtCustomClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}

// 自定义验证器
type CustomValidator struct {
	once      sync.Once
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.lazyInit()
	if err := cv.validator.Struct(i); err != nil {
		// 返回验证错误
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// 初始化验证器
func (cv *CustomValidator) lazyInit() {
	cv.once.Do(func() {
		cv.validator = validator.New()
	})
}

func main() {
	e := echo.New()

	// 使用自定义验证器
	e.Validator = &CustomValidator{}

	// 中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	// 健康检查
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"pid":    fmt.Sprint(os.Getpid()),
			"status": "OK",
		})
	})

	// API 路由分组

	// 分享 API 路由分组
	share_v1 := e.Group("/api/v1/share")
	{
		share_v1.POST("/text", share.Text)
		share_v1.GET("/text", share.Text)
		share_v1.POST("/file", share.File)
		share_v1.GET("/file", share.File)
		share_v1.POST("/select", share.Select)
		share_v1.POST("/download", share.Download)
		share_v1.GET("/code/get", share.GetCodeFile)
	}

	// 管理 API 路由分组
	// 使用JWT中间件验证管理员身份

	admin_jwt := echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte("secret"),
		TokenLookup: "header:Authorization",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
	})

	admin_v1 := e.Group("/api/v1/admin")
	admin_v1.Use(admin_jwt)
	{
		admin_v1.POST("/login", func(c echo.Context) error {
			// 设置自定义claims
			claims := &JwtCustomClaims{
				Username: "admin",
				Admin:    true,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				},
			}

			// 创建token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// 签名token
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		})
		admin_v1.GET("/dashboard", admin.Dashboard)
		admin_v1.DELETE("/file/delete", admin.DeleteFile)
		admin_v1.GET("/file/list", admin.FileList)
		admin_v1.GET("/config/get", admin.GetConfig)
		admin_v1.POST("/config/update", admin.UpdateConfig)
		admin_v1.POST("/file/download", admin.FileDownload)
		admin_v1.GET("/local/list", admin.GetLocalList)
		admin_v1.DELETE("/local/delete", admin.DeleteLocalFile)
		admin_v1.POST("/local/share", admin.ShareLocalFile)
		admin_v1.GET("/me", func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*JwtCustomClaims)
			return c.JSON(http.StatusOK, claims)
		})
	}

	// 服务器配置
	server := &http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	// 分组列出所有路由
	e.GET("/api/v1/routes", func(c echo.Context) error {
		// 根据path排序
		routes := e.Routes()
		sort.Slice(routes, func(i, j int) bool {
			return routes[i].Path < routes[j].Path
		})
		for _, route := range routes {
			if route.Path == "/api/v1/routes" {
				continue
			}
			fmt.Println(route.Path, route.Method)
		}
		return c.JSON(http.StatusOK, routes)
	})

	// 启动服务
	e.Logger.Fatal(e.StartServer(server))
}
