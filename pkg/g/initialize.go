package g

import (
	"path"
	"runtime"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin-frame/pkg/cfg"
	"gin-frame/pkg/gorm"
	"gin-frame/pkg/log"
)

func InitCfg(file string) error {
	// 是否使用相对路径
	if file[:1] == "." {
		_, filename, _, _ := runtime.Caller(1)     // 获取运行文件路径
		file = path.Join(path.Dir(filename), file) // 拼接完整的配置文件路径
	}

	return cfg.Init(file)
}

func InitDB() error {
	return gorm.InitDB()
}

func InitLog() {
	log.Init()
}
