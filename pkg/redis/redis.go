package redis

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"

	"frame/pkg/cfg"
)

var (
	once   sync.Once
	DB     *redis.Client
	config *Config
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func Init() {
	DB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
	})
}

func InitConfig() {
	if err := cfg.Viper().UnmarshalKey("redis", &config); err != nil {
		panic(fmt.Errorf("初始化Redis配置错误 err: %v", err))
	}
}

func Redis() *redis.Client {
	once.Do(func() {
		InitConfig()
		Init()
	})

	return DB
}
