package user

import (
	"github.com/gin-gonic/gin"

	s_user "frame/internal/services/user"
	"frame/pkg/errno"
	"frame/pkg/g"
	"frame/pkg/log"
	"frame/pkg/response"
)

func Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		g.Log().Errorf("register bind param err: %v", err)
		response.Send(ctx, errno.ErrBind, nil)
		ctx.Abort()
		return
	}

	// 验证参数
	if g.IsEmpty(req.Username) || g.IsEmpty(req.Password) || g.IsEmpty(req.Email) {
		response.Send(ctx, errno.ErrParam, nil)
		return
	}
	log.Debug("email=", req.Email)

	// 验证重复密码
	if req.Password != req.ConfirmPassword {
		response.Send(ctx, errno.ErrTwicePasswordNotMatch, nil)
		return
	}

	// 执行注册
	if err := s_user.Svc.Register(ctx, req.Username, req.Email, req.Password); err != nil {
		response.Send(ctx, errno.ErrRegisterFailed, nil)
		return
	}

	response.Send(ctx, errno.OK, nil)
}
