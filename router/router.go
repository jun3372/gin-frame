package router

import (
	"github.com/gin-gonic/gin"

	"frame/internal/handlers/v1/user"
	"frame/internal/middleware"
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
	v1 := router.Group("/v1")
	{
		v1.POST("/register", user.Register)
		v1.POST("/login/phone", user.PhoneLogin)
		v1.POST("/login", user.EmailLogin)

		// 分组用户, 并且需要授权
		u := v1.Group("/users")
		u.Use(middleware.Auth())
		{
			u.PUT("/", user.Update)
			// u.POST("/follow", user.Follow)
			// u.GET("/:id/following", user.FollowList)
			// u.GET("/:id/followers", user.FollowerList)
		}
	}
	// router.POST("/v1/login", user.Login)
	// router.POST("/v1/login/phone", user.PhoneLogin)
	// router.GET("/v1/vcode", user.VCode)
}
