package s_user

import (
	"github.com/gin-gonic/gin"

	muser "frame/internal/models/user"
	"frame/internal/repository/user"
	"frame/pkg/auth"
)

type Service interface {
	Register(ctx *gin.Context, username, email, password string) error
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
