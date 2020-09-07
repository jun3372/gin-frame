package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jun3372/gin-frame/internal/middleware"
	// muser "github.com/jun3372/gin-frame/internal/models/user"
	"github.com/jun3372/gin-frame/pkg/errno"
	"github.com/jun3372/gin-frame/pkg/frame"
	"github.com/jun3372/gin-frame/pkg/response"
	"github.com/jun3372/gin-frame/router"
)

func main() {
	// 实例化应用
	app := frame.New()

	// 注册数据库迁移结构体
	// app.AutoMigrate(&muser.UserBaseModel{}, &muser.UserFansModel{}, &muser.UserFollowModel{})

	// 加载更多路由
	router.Load(app.Router)

	// 注册路由
	app.AddRouter(http.MethodGet, "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"content": "请求成功拉"})
	})

	u := app.Router.Group("/v1/users")
	{
		u.Use(middleware.Auth())
		u.GET("/:id", func(c *gin.Context) {
			response.Send(c, errno.OK, gin.H{"id": c.Param("id")})
		})
	}

	app.Run()
}
