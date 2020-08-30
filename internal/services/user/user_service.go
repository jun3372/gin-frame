package s_user

import (
	"github.com/gin-gonic/gin"

	muser "frame/internal/models/user"
	"frame/internal/repository/user"
	"frame/pkg/auth"
	"frame/pkg/errno"
	"frame/pkg/token"
)

type Service interface {
	Register(ctx *gin.Context, username, email, password string) error
	EmailLogin(ctx *gin.Context, email, password string) (tokenStr string, err error)
}

type userService struct {
	userRepo user.BaseRepo
}

// Svc 直接初始化，可以避免在使用时再实例化
var Svc = NewUserService()

func NewUserService() *userService {
	return &userService{
		userRepo: user.NewUserRepo(),
	}
}

func (srv *userService) Register(ctx *gin.Context, username, email, password string) (err error) {
	pwd, err := auth.Encrypt(password)
	if err != nil {
		return
	}

	model := muser.UserBaseModel{
		Username: username,
		Password: pwd,
		Phone:    0,
		Email:    email,
		Avatar:   "",
		Sex:      0,
	}

	_, err = srv.userRepo.Create(model)
	return
}

func (srv *userService) EmailLogin(ctx *gin.Context, email, password string) (tokenStr string, err error) {
	user, err := srv.userRepo.GetUserByEmail(email)
	if err != nil {
		return
	}

	// 验证密码
	if err = auth.Compare(user.Password, password); err != nil {
		err = errno.ErrPasswordIncorrect
		return
	}

	// 执行签名
	tokenStr, err = token.Sign(ctx, token.Context{
		UserID:   user.ID,
		Username: user.Username,
	}, "")

	return
}
