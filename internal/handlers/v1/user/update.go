package user

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	suser "frame/internal/services/user"
	"frame/pkg/errno"
	"frame/pkg/g"
	"frame/pkg/response"
)

func Update(ctx *gin.Context) {
	var (
		userId uint64
		err    error
		req    UpdateRequest
		params g.Map
	)

	userId = cast.ToUint64(ctx.Param("id"))
	if err = ctx.Bind(&req); err != nil {
		response.Send(ctx, errno.ErrBind, nil)
		return
	}

	if g.IsEmpty(userId) {
		response.Send(ctx, errno.ErrParam, nil)
		return
	}

	params = g.Map{"avatar": req.Avatar, "sex": req.Sex}
	if err = suser.Svc.UpdateUser(userId, params); err != nil {
		response.Send(ctx, err, nil)
		return
	}

	response.Send(ctx, errno.OK, nil)
}
