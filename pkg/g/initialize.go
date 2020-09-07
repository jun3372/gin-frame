package g

import (
	"path"
	"runtime"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jun3372/gin-frame/pkg/cfg"
	"github.com/jun3372/gin-frame/pkg/gorm"
	"github.com/jun3372/gin-frame/pkg/log"
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
