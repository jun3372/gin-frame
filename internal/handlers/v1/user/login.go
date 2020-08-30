package user

import (
	"github.com/gin-gonic/gin"

	muser "frame/internal/models/user"
	s_user "frame/internal/services/user"
	"frame/pkg/errno"
	"frame/pkg/g"
	"frame/pkg/response"
)

func Login(ctx *gin.Context) {
	var req LoginCredentials
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Send(ctx, errno.ErrBind, nil)
		return
	}

	if g.IsEmpty(req.Email) || g.IsEmpty(req.Password) {
		response.Send(ctx, errno.ErrParam, nil)
		return
	}

	token, err := s_user.Svc.EmailLogin(ctx, req.Email, req.Password)
	if err != nil {
		response.Send(ctx, err, nil)
		return
	}

	response.Send(ctx, nil, muser.Token{Token: token})
}
