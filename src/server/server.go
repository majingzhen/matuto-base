package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/global"
	"matuto-base/src/routers"
	"matuto-base/src/utils"
	"net/http"
	"os"
	"os/signal"
)

func init() {
	global.InitViper()
	global.InitLogger()
	global.InitDataSource()
	// global.InitMongoClient()
}

func Run() {
	gin.SetMode(global.Viper.GetString("server.model"))
	// 关闭日志颜色
	gin.DisableConsoleColor()
	router := new(routers.Routers).InitRouter()
	port := global.Viper.GetInt("server.port")
	// 创建一个HTTP服务
	// 启动服务
	go func() {
		if err := router.Run(fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Fatal("Server startup failed", zap.Error(err))
		}
	}()
	global.Logger.Info(fmt.Sprintf("系统启动成功,服务运行在 http://%s:%d", utils.GetIp(), port))
	// 创建一个信号接收器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	// 等待接收到关闭信号
	<-quit
	global.Logger.Info("Shutting down server...")
	global.Logger.Info("Server exiting")

}
