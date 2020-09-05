package s_user

import (
	"github.com/gin-gonic/gin"

	muser "gin-frame/internal/models/user"
	"gin-frame/internal/repository/user"
	"gin-frame/pkg/auth"
	"gin-frame/pkg/errno"
	"gin-frame/pkg/g"
	"gin-frame/pkg/token"
)

type Service interface {
	Register(ctx *gin.Context, username, email, password string) error
	EmailLogin(ctx *gin.Context, email, password string) (tokenStr string, err error)
	PhoneLogin(ctx *gin.Context, phone, verifyCode int) (tokenStr string, err error)
	GetUserByID(id uint64) (*muser.UserBaseModel, error)
	GetUserInfoByID(id uint64) (*muser.UserInfo, error)
	GetUserByPhone(phone int) (*muser.UserBaseModel, error)
	GetUserByEmail(email string) (*muser.UserBaseModel, error)
	UpdateUser(id uint64, userMap g.Map) error
	BatchGetUsers(userID uint64, userIDs []uint64) ([]*muser.UserInfo, error)

	// 关注
	IsFollowedUser(userID uint64, followedUID uint64) bool
	AddUserFollow(userID uint64, followedUID uint64) error
	CancelUserFollow(userID uint64, followedUID uint64) error
	GetFollowingUserList(userID uint64, lastID uint64, limit int) ([]*muser.UserFollowModel, error)
	GetFollowerUserList(userID uint64, lastID uint64, limit int) ([]*muser.UserFansModel, error)
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
		Email:    email,
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

func (srv *userService) PhoneLogin(ctx *gin.Context, phone, verifyCode int) (tokenStr string, err error) {
	// TODO:: 验证码手机验证码

	// 获取用户信息
	user, err := srv.userRepo.GetUserByPhone(phone)
	if err != nil {
		return
	}

	// 签名
	tokenStr, err = token.Sign(ctx, token.Context{UserID: user.ID, Username: user.Username}, "")
	return
}

func (srv *userService) GetUserByID(id uint64) (*muser.UserBaseModel, error) {
	panic("implement me")
}

func (srv *userService) GetUserInfoByID(id uint64) (*muser.UserInfo, error) {
	panic("implement me")
}

func (srv *userService) GetUserByPhone(phone int) (*muser.UserBaseModel, error) {
	panic("implement me")
}

func (srv *userService) GetUserByEmail(email string) (*muser.UserBaseModel, error) {
	panic("implement me")
}

func (srv *userService) UpdateUser(id uint64, userMap g.Map) error {
	if err:=  srv.userRepo.Update(id, userMap);err!=nil{
		return err
	}

	return nil
}

func (srv *userService) BatchGetUsers(userID uint64, userIDs []uint64) ([]*muser.UserInfo, error) {
	panic("implement me")
}

func (srv *userService) IsFollowedUser(userID uint64, followedUID uint64) bool {
	panic("implement me")
}

func (srv *userService) AddUserFollow(userID uint64, followedUID uint64) error {
	panic("implement me")
}

func (srv *userService) CancelUserFollow(userID uint64, followedUID uint64) error {
	panic("implement me")
}

func (srv *userService) GetFollowingUserList(userID uint64, lastID uint64, limit int) ([]*muser.UserFollowModel, error) {
	panic("implement me")
}

func (srv *userService) GetFollowerUserList(userID uint64, lastID uint64, limit int) ([]*muser.UserFansModel, error) {
	panic("implement me")
}
