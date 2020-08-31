package user

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cast"

	muser "frame/internal/models/user"
	"frame/pkg/g"
	"frame/pkg/log"
)

// BaseRepo 定义用户仓库接口
type BaseRepo interface {
	Create(user muser.UserBaseModel) (id uint64, err error)
	Update(id uint64, userMap g.Map) error
	GetUserByID(id uint64) (*muser.UserBaseModel, error)
	GetUsersByIds(ids []uint64) ([]*muser.UserBaseModel, error)
	GetUserByPhone(phone int) (*muser.UserBaseModel, error)
	GetUserByEmail(email string) (*muser.UserBaseModel, error)
}

// userRepo 用户仓库
type userRepo struct {
	// userCache *user.Cache
}

func (repo *userRepo) Update(id uint64, userMap g.Map) error {
	return g.DB().Model(muser.UserBaseModel{ID: id}).Updates(userMap).Error
}

func (repo *userRepo) GetUserByID(id uint64) (*muser.UserBaseModel, error) {
	panic("implement me")
}

func (repo *userRepo) GetUsersByIds(ids []uint64) ([]*muser.UserBaseModel, error) {
	panic("implement me")
}

// Create 创建用户
func (repo *userRepo) Create(user muser.UserBaseModel) (id uint64, err error) {
	j, _ := json.Marshal(user)
	log.Debug("user=", cast.ToString(j))
	if err = g.DB().Create(&user).Error; err != nil {
		return 0, errors.Wrap(err, "[user_repo] create user err")
	}

	return user.ID, nil
}

// GetUserByPhone 根据手机号获取用户
func (repo *userRepo) GetUserByPhone(phone int) (*muser.UserBaseModel, error) {
	user := muser.UserBaseModel{}
	err := g.DB().Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "[user_repo] get user err by phone")
	}

	return &user, nil
}

// GetUserByEmail 根据邮箱获取手机号
func (repo *userRepo) GetUserByEmail(phone string) (*muser.UserBaseModel, error) {
	user := muser.UserBaseModel{}
	err := g.DB().Where("email = ?", phone).First(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "[user_repo] get user err by email")
	}

	return &user, nil
}

// NewUserRepo 实例化用户仓库
func NewUserRepo() BaseRepo {
	return &userRepo{}
}
