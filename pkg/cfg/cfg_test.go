package cfg

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInit(t *testing.T) {
	var err error
	if err = Init("../../config/config.yml"); err != nil {
		t.Fatal(err)
	}

	// database := gorm.DbConfig{
	// 	Type:     "mysql",
	// 	Host:     "127.0.0.1",
	// 	Port:     3306,
	// 	User:     "root",
	// 	Password: "root",
	// 	Dbname:   "test",
	// 	Charset:  "utf-8",
	// 	Sslmode:  "disable",
	// 	Link:     "",
	// }
	// t.Logf("database=%v", database)
	// viper.Set("database", database)

	// config := log.Config{
	// 	Name:      "渣渣辉",
	// 	Path:      "./runtime/log",
	// 	Level:     "",
	// 	Formatter: "json",
	// 	Stdout:    true,
	// }
	// viper.Set("logger", config)


	if err = viper.WriteConfig(); err != nil {
		panic(err)
	}
}
