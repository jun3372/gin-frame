package user

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	suser "github.com/jun3372/gin-frame/internal/services/user"
	"github.com/jun3372/gin-frame/pkg/errno"
	"github.com/jun3372/gin-frame/pkg/g"
	"github.com/jun3372/gin-frame/pkg/response"
)

func Update(ctx *gin.Context) {
	var (
		userId int64
		err    error
		req    UpdateRequest
		params g.Map
	)

	userId = ctx.GetInt64(response.GetUserIdKey())
	if g.IsEmpty(userId) {
		response.Send(ctx, errno.ErrParam, nil)
		return
	}

	if err = ctx.Bind(&req); err != nil {
		response.Send(ctx, errno.ErrBind, nil)
		return
	}

	params = g.Map{"avatar": req.Avatar, "sex": req.Sex}
	if err = suser.Svc.UpdateUser(cast.ToUint64(userId), params); err != nil {
		response.Send(ctx, err, nil)
		return
	}

	response.Send(ctx, errno.OK, nil)
}
