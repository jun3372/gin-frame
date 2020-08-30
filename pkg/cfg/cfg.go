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
	c *viper.Viper
)

func Init(file string) (err error) {
	if empty.IsEmpty(file) {
		flag.StringVar(&file, "cfg", "", "Configuration file path")
		flag.Parse()

		// 是否跳过初始化配置文件
		if empty.IsEmpty(file) {
			panic(fmt.Errorf("读取输入配置文件错误: file=%v", file))
		}

		// 是否使用相对路径
		if file[:1] == "." {
			dir, _ := os.Getwd()        // 获取当前运行目录
			file = path.Join(dir, file) // 拼接完整的配置文件路径
		}
	}

	if empty.IsEmpty(file) {
		panic(fmt.Errorf("读取配置文件错误: file=%v", file))
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

	c = viper.GetViper()
	return nil
}

func Viper() *viper.Viper {
	if empty.IsEmpty(c) || empty.IsNil(c) {
		_ = Init("")
	}

	return c
}
