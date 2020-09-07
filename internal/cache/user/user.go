package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	muser "github.com/jun3372/gin-frame/internal/models/user"
	redis2 "github.com/jun3372/gin-frame/pkg/redis"
)

const (
	// PrefixCacheKey 业务cache key
	PrefixCacheKey = "xxxx"
	// PrefixUserBaseCacheKey cache前缀
	PrefixUserBaseCacheKey = "user:cache:%d"
	// DefaultExpireTime 默认过期时间
	DefaultExpireTime = time.Hour * 24
)

type Cache struct {
	cache *redis.Client
}

func NewUserCache() *Cache {
	return &Cache{cache: redis2.Redis()}
}

// GetUserBaseCacheKey 获取cache key
func (u *Cache) GetUserBaseCacheKey(userID uint64) string {
	return fmt.Sprintf(PrefixCacheKey+":"+PrefixUserBaseCacheKey, userID)
}

// SetUserBaseCache 写入用户cache
func (u *Cache) SetUserBaseCache(userID uint64, user *muser.UserBaseModel) error {
	if user == nil || user.ID == 0 {
		return nil
	}

	cacheKey := fmt.Sprintf(PrefixUserBaseCacheKey, userID)
	cache, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := u.cache.Set(context.TODO(), cacheKey, cache, DefaultExpireTime).Err(); err != nil {
		return err
	}

	return nil
}

// GetUserBaseCache 获取用户cache
func (u *Cache) GetUserBaseCache(userID uint64) (userModel *muser.UserBaseModel, err error) {
	cacheKey := fmt.Sprintf(PrefixUserBaseCacheKey, userID)
	user := u.cache.Get(context.TODO(), cacheKey)
	if err = user.Err(); err != nil {
		return
	}

	// err = user.Scan(userModel)
	err = json.Unmarshal([]byte(user.Val()), &userModel)
	return
}

// MultiGetUserBaseCache 批量获取用户cache
func (u *Cache) MultiGetUserBaseCache(userIDs []uint64) (users []*muser.UserBaseModel, err error) {
	var cacheKeys []string
	for _, v := range userIDs {
		cacheKey := fmt.Sprintf(PrefixUserBaseCacheKey, v)
		cacheKey, err := BuildCacheKey(PrefixCacheKey, cacheKey)
		if err != nil {
			return nil, errors.Wrapf(err, "build cache key err, key is %+v", cacheKey)
		}

		cacheKeys = append(cacheKeys, cacheKey)
	}

	values, err := u.cache.MGet(context.TODO(), cacheKeys...).Result()
	if err != nil {
		return nil, err
	}

	users = []*muser.UserBaseModel{}
	for _, value := range values {
		bytes := []byte(cast.ToString(value))
		var user *muser.UserBaseModel
		if err = json.Unmarshal(bytes, &user); err != nil {
			fmt.Println("err=", err)
		}

		users = append(users, user)
	}

	return
}

// DelUserBaseCache 删除用户cache
func (u *Cache) DelUserBaseCache(userID uint64) error {
	cacheKey := fmt.Sprintf(PrefixUserBaseCacheKey, userID)
	err := u.cache.Del(context.TODO(), cacheKey).Err()
	if err != nil {
		return err
	}

	return nil
}

// BuildCacheKey 构建一个带有前缀的缓存key
func BuildCacheKey(keyPrefix string, key string) (cacheKey string, err error) {
	if key == "" {
		return "", errors.New("[cache] key should not be empty")
	}

	cacheKey, err = strings.Join([]string{keyPrefix, key}, ":"), nil
	return
}
