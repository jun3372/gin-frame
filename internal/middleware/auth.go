package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gin-frame/pkg/errno"
	"gin-frame/pkg/log"
	"gin-frame/pkg/response"
	"gin-frame/pkg/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		log.Infof("context is: %+v", ctx)

		if err != nil {
			response.Send(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		// set uid to context
		c.Set(response.GetUserIdKey(), cast.ToInt64(ctx.UserID))
		c.Next()
	}
}
