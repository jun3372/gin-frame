package gorm

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"frame/pkg/cfg"
	"frame/pkg/empty"
	"frame/pkg/log"
)

type Config struct {
	Type     string `json:"type"` // 类型: MySQL, PostgreSQL, Sqlite3
	Host     string `json:"host"` // 服务器地址
	Port     int    `json:"port"`
	User     string `json:"user"`     // 链接用户
	Password string `json:"password"` // 链接密码
	Dbname   string `json:"dbname"`   // 数据库名
	Charset  string `json:"charset"`  // 数据库字符集
	Sslmode  string `json:"sslmode"`  // PostgreSQL的sslmode: disable
	Link     string `json:"link"`     // 链接数据库
	Debug    bool   `json:"debug"`    // 是否开启日志
}

var (
	db     *gorm.DB
	once   sync.Once
	config *Config
)

// 初始化数据库配置
func InitConfig() (*Config, error) {
	if err := cfg.Viper().UnmarshalKey("database", &config); err != nil {
		return config, err
	}

	return config, nil
}

// 初始化数据库链接
func InitDB() (err error) {
	if empty.IsEmpty(config) {
		if _, err := InitConfig(); err != nil {
			panic(fmt.Errorf("初始化数据库配置错误, err: %v", err))
		}
	}

	var (
		t    = config.GetType("mysql")
		link = config.GetLink()
	)

	db, err = gorm.Open(t, link)
	db.SetLogger(log.Logger)
	db.LogMode(config.Debug)
	return err
}

// 关闭数据库链接
func CloseDB() error {
	return db.Close()
}

// 获取数据库链接
func GetDB() *gorm.DB {
	once.Do(func() {
		_ = InitDB()
	})

	return db
}

func (c *Config) GetLink() string {
	switch c.GetType("mysql") {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Dbname, c.Charset)
	case "postgres":
		return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", c.Host, c.User, c.Dbname, c.Sslmode, c.Password)
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", c.User, c.Password, c.Host, c.Port, c.Dbname)

	}
	return ""
}

func (c *Config) GetType(def string) string {
	if empty.IsEmpty(c.Type) {
		return def
	}

	return strings.ToLower(c.Type)
}
