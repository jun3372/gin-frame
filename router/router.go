package router

import (
	"github.com/gin-gonic/gin"

	"frame/internal/handlers/v1/user"
	"frame/pkg/response"
)

func Load(router *gin.Engine, handlerFunc ...gin.HandlerFunc) {
	// router.Use(middleware.Options)
	router.Use(handlerFunc...)

	// 404 Handler.
	router.NoRoute(response.RouteNotFound)
	router.NoMethod(response.RouteNotFound)

	// 静态资源，主要是图片
	// router.Static("/static", "./static")

	// 认证相关路由
	router.POST("/v1/register", user.Register)
	// router.POST("/v1/login", user.Login)
	// router.POST("/v1/login/phone", user.PhoneLogin)
	// router.GET("/v1/vcode", user.VCode)
}
