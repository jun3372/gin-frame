package user

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"gin-frame/internal/cache/user"
	muser "gin-frame/internal/models/user"
	"gin-frame/pkg/errno"
	"gin-frame/pkg/g"
	"gin-frame/pkg/log"
	"gin-frame/pkg/redis"
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
	userCache *user.Cache
}

func (repo *userRepo) Update(id uint64, userMap g.Map) error {
	user, err := repo.GetUserByID(id)
	if err != nil {
		return err
	}

	if g.IsEmpty(user) || user.ID < 1 {
		return errno.ErrUserNotFound
	}

	// 删除缓存
	if err := repo.userCache.DelUserBaseCache(id); err != nil {
		return err
	}

	return g.DB().Model(user).Updates(userMap).Error
}

func (repo *userRepo) GetUserByID(id uint64) (*muser.UserBaseModel, error) {
	userModel, err := repo.userCache.GetUserBaseCache(id)
	if err == nil && !g.IsEmpty(userModel) && userModel.ID > 0 {
		return userModel, nil
	}

	// 加锁，防止缓存击穿
	key := fmt.Sprintf("uid:%d", id)
	lock := redis.NewLock(g.Redis(), key, 3*time.Second)
	token := lock.GenToken()
	isLock, err := lock.Lock(token)
	if err != nil || !isLock {
		return nil, errors.Wrap(err, "[user_repo] lock err")
	}

	// 解除锁定
	defer func() {
		_ = lock.Unlock(token)
	}()

	data := muser.UserBaseModel{}
	if isLock {
		// 从数据库中获取
		err = g.DB().Where("`id` = ?", id).First(&data).Error
		log.Debug("err=", err)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errors.Wrap(err, "[user_repo] get user data err")
		}

		// 写入cache
		err = repo.userCache.SetUserBaseCache(id, &data)
		if err != nil {
			return &data, errors.Wrap(err, "[user_repo] set user data err")
		}
	}

	return &data, nil
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
	err := g.DB().Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "[user_repo] get user err by email")
	}

	return &user, nil
}

// NewUserRepo 实例化用户仓库
func NewUserRepo() BaseRepo {
	return &userRepo{userCache: user.NewUserCache()}
}
