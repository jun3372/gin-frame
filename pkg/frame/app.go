package frame

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jun3372/gin-frame/internal/middleware"
	"github.com/jun3372/gin-frame/pkg/cfg"
	"github.com/jun3372/gin-frame/pkg/gorm"
	"github.com/jun3372/gin-frame/pkg/log"
)

const (
	// ModeDebug debug mode
	ModeDebug string = "debug"
	// ModeRelease release mode
	ModeRelease string = "release"
	// ModeTest test mode
	ModeTest string = "test"
)

// App is singleton
var App *Application

type Application struct {
	Router *gin.Engine
	Debug  bool
}

// New Application
func New() *Application {
	App = new(Application)
	switch strings.ToLower(cfg.Viper().GetString("app.mode")) {
	case ModeRelease:
		gin.SetMode(ModeRelease)
		break
	case ModeTest:
		gin.SetMode(ModeTest)
		break
	case ModeDebug:
	default:
		gin.SetMode(ModeDebug)
	}

	App.println()

	// init router
	App.Router = gin.Default()

	// 全局中间件
	App.Router.Use(middleware.Logger(), middleware.RequestID(), middleware.Jaeger(), middleware.Options, gin.Recovery())

	// 获取配置文件信息
	if cfg.Viper().GetString("app.mode") == ModeDebug {
		App.Debug = true
	}

	return App
}

// Run start a app
func (a *Application) Run() {
	// log.Infof("Start to listening the incoming requests on http address: %s", cfg.Viper().GetString("app.addr"))
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Viper().GetString("app.addr"), cfg.Viper().GetString("app.port")),
		Handler: a.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err.Error())
		}
	}()

	// 优雅退出
	gracefulStop(srv)
}

// 注册中间件
func (a *Application) AddMiddleware(middleware ...gin.HandlerFunc) {
	a.Router.Use(middleware...)
}

// 添加路由
func (a *Application) AddRouter(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	a.Router.Handle(httpMethod, relativePath, handlers...)
}

// 注册数据库迁移结构体
func (a *Application) AutoMigrate(m ...interface{}) {
	gorm.GetDB().AutoMigrate(m...)
}

func (a *Application) println() {
	fmt.Println("")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|            go-gin-frame           |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port: " + cfg.Viper().GetString("app.port") + "     Pid: " + fmt.Sprintf("%d", os.Getpid()) + "      |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")
}

// gracefulStop 优雅退出
// 等待中断信号以超时 5 秒正常关闭服务器
// 官方说明：https://github.com/gin-gonic/gin#graceful-restart-or-stop
func gracefulStop(srv *http.Server) {
	quit := make(chan os.Signal)
	// kill 命令发送信号 syscall.SIGTERM
	// kill -2 命令发送信号 syscall.SIGINT
	// kill -9 命令发送信号 syscall.SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// 5 秒后捕获 ctx.Done() 信号
	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	default:
	}

	log.Info("Server exiting")
}
