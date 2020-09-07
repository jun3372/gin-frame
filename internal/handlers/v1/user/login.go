package user

import (
	"github.com/gin-gonic/gin"

	muser "frame/internal/models/user"
	suser "frame/internal/services/user"
	"frame/pkg/errno"
	"frame/pkg/g"
	"frame/pkg/log"
	"frame/pkg/response"
)

func EmailLogin(ctx *gin.Context) {
	var (
		err   error
		token string
		req   LoginCredentials
	)

	if err = ctx.Bind(&req); err != nil {
		response.Send(ctx, errno.ErrBind, nil)
		return
	}

	if g.IsEmpty(req.Email) || g.IsEmpty(req.Password) {
		response.Send(ctx, errno.ErrParam, nil)
		return
	}

	if token, err = suser.Svc.EmailLogin(ctx, req.Email, req.Password); err != nil {
		response.Send(ctx, err, nil)
		return
	}

	response.Send(ctx, nil, muser.Token{Token: token})
}

func PhoneLogin(ctx *gin.Context) {
	var (
		err  error
		req  PhoneLoginCredentials
	)

	// 获取绑定参数
	if err = ctx.Bind(&req); err != nil {
		log.Debug("绑定参数错误", req)
		response.Send(ctx, errno.ErrBind, nil)
		return
	}

	// 验证绑定参数
	if g.IsEmpty(req.Phone) || g.IsEmpty(req.VerifyCode) {
		log.Debug("参数错误", req)
		response.Send(ctx, errno.ErrParam, nil)
		return
	}

	// 登录
	token, err := suser.Svc.PhoneLogin(ctx, req.Phone, req.VerifyCode)
	if err != nil {
		response.Send(ctx, err, nil)
		return
	}

	response.Send(ctx, nil, muser.Token{Token: token})
}
