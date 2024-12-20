package routers

import (
	"github.com/gin-gonic/gin"
	genRouter "matuto-base/src/app/admin/gen/router"
	sysRouter "matuto-base/src/app/admin/sys/router"
	"matuto-base/src/middleware"
)

type Routers struct {
	baseRouter BaseRouter
	sysRouter  sysRouter.SysRouter
	genRouter  genRouter.GenRouter
}

// InitRouter 初始化路由
func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(gin.Recovery())

	// 跨域处理
	// 使用Cors中间件处理跨域请求
	r.Use(middleware.Cors())
	api := r.Group("/api")
	{
		routers.sysRouter.InitSysRouter(api)
		routers.baseRouter.InitBaseRouter(api)
		routers.genRouter.InitGenRouter(api)
	}

	return r
}
