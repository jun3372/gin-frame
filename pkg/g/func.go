package g

import (
	"fmt"
	"runtime"

	redis2 "github.com/go-redis/redis/v8"
	gorm2 "github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gin-frame/pkg/cfg"
	"gin-frame/pkg/empty"
	"gin-frame/pkg/gorm"
	"gin-frame/pkg/log"
	"gin-frame/pkg/redis"
)

func IsNil(value interface{}) bool {
	return empty.IsNil(value)
}

func IsEmpty(value interface{}) bool {
	return empty.IsEmpty(value)
}

func DB() *gorm2.DB {
	return gorm.GetDB()
}

func Log() *logrus.Entry {
	_, file, line, _ := runtime.Caller(1)
	caller := fmt.Sprintf("%s:%d", file, line)
	return log.GetLog().WithField("caller", caller)
}

func Config() *viper.Viper {
	return cfg.Viper()
}

func Redis() *redis2.Client {
	return redis.Redis()
}
