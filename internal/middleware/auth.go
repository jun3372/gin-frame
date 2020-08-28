package middleware

import (
	"github.com/gin-gonic/gin"

	"frame/pkg/errno"
	"frame/pkg/log"
	"frame/pkg/response"
	"frame/pkg/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		log.Log().Infof("context is: %+v", ctx)

		if err != nil {
			response.Send(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		// set uid to context
		c.Set(response.GetUserKey(), ctx.UserID)
		c.Next()
	}
}
