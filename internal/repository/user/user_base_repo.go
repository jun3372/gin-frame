package user

import (
	"github.com/pkg/errors"

	muser "frame/internal/models/user"
	"frame/pkg/g"
)

// BaseRepo 定义用户仓库接口
type BaseRepo interface {
	Create(user muser.UserBaseModel) (id uint64, err error)
	// Update(db *gorm.DB, id uint64, userMap map[string]interface{}) error
	// GetUserByID(db *gorm.DB, id uint64) (*muser.UserBaseModel, error)
	// GetUsersByIds(db *gorm.DB, ids []uint64) ([]*muser.UserBaseModel, error)
	GetUserByPhone(phone int) (*muser.UserBaseModel, error)
	GetUserByEmail(email string) (*muser.UserBaseModel, error)
}

// userRepo 用户仓库
type userRepo struct {
	// userCache *user.Cache
}

// NewUserRepo 实例化用户仓库
func NewUserRepo() BaseRepo {
	return &userRepo{}
}

// Create 创建用户
func (repo *userRepo) Create(user muser.UserBaseModel) (id uint64, err error) {
	err = g.DB().Create(&user).Error
	if err != nil {
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
