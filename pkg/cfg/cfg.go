package cfg

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"

	"frame/pkg/empty"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "cfg", "./config.yml", "Configuration file path")
	flag.Parse()

	// 是否使用相对路径
	if file[:1] == "." {
		dir, _ := os.Getwd()        // 获取当前运行目录
		file = path.Join(dir, file) // 拼接完整的配置文件路径
	}

	if err := Init(file); err != nil {
		panic(fmt.Errorf("Failed to load configuration file, err: %v", err))
	}
}

func Init(file string) (err error) {
	if empty.IsEmpty(file) {
		panic(fmt.Errorf("输入配置文件错误: file=%v", file))
	}

	// 是否使用相对路径
	if file[:1] == "." {
		_, filename, _, _ := runtime.Caller(1)     // 获取运行文件路径
		file = path.Join(path.Dir(filename), file) // 拼接完整的配置文件路径
	}

	// 设置需要读取的配置文件
	viper.SetConfigFile(file)
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return nil
}

func Viper() *viper.Viper {
	return viper.GetViper()
}
