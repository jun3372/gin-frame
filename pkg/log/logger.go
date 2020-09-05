package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"gin-frame/pkg/cfg"
	"gin-frame/pkg/empty"
)

type Config struct {
	Name      string `json:"name"`      // 项目名称: xxx
	Path      string `json:"path"`      // 日志路径: ./
	Level     string `json:"level"`     // 日志等级: 默认: TraceLevel
	Formatter string `json:"formatter"` // 日志格式: json, text 默认: text
	Stdout    bool   `json:"stdout"`    // 是否控制台输出
}

var (
	Once sync.Once

	config *Config
	Logger *logrus.Logger
)

func SetConfig(c *Config) {
	config = c
}

func init() {
	if empty.IsEmpty(Logger) || empty.IsNil(Logger) {
		Logger = logrus.New()
	}
}

func InitConfig() {
	if !empty.IsEmpty(config) && !empty.IsNil(config) {
		return
	}

	if err := cfg.Viper().UnmarshalKey("logger", &config); err != nil {
		panic(fmt.Errorf("初始化日志配置错误 err: %v", err))
	}
}

func Init() {
	InitConfig()

	// 设置日志格式
	if config.Formatter == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{})
	}

	if config.Stdout {
		// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
		// 日志消息输出可以是任意的io.writer类型
		Logger.SetOutput(os.Stdout)
	}

	// 设置日志级别
	switch strings.ToLower(config.Level) {
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
		break
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
		break
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
		break
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
		break
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
		break
	case "trace":
	default:
		Logger.SetLevel(logrus.TraceLevel)
	}

	// Logger = Logger.WithField("caller", caller).Logger

	// 添加Hook
	if hook, err := newHook(); err != nil {
		panic(fmt.Errorf("config local file system for logger error: %v", err))
	} else {
		Logger.AddHook(hook)
	}
}

func newHook() (*lfshook.LfsHook, error) {
	fileName := fmt.Sprintf("%s/%s", config.Path, config.Name)
	writer, err := rotatelogs.New(
		fileName+"-%Y-%m-%d.log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(fileName),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Hour*24*10),
	)

	if err != nil {
		return nil, err
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
	return lfsHook, err
}

func GetLog() *logrus.Entry {
	Once.Do(func() {
		Init()
	})

	_, file, line, _ := runtime.Caller(1)
	caller := fmt.Sprintf("%s:%d", file, line)
	return Logger.WithField("caller", caller)
}
