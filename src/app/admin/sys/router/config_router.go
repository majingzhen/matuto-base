// Package router ConfigRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type ConfigRouter struct {
	configApi api.ConfigApi
}

// InitConfigRouter 初始化 Config 路由信息
func (r *ConfigRouter) InitConfigRouter(Router *gin.RouterGroup) {
	configRouter := Router.Group("config").Use(middleware.JWTAuthFilter())
	configRouterWithoutRecord := Router.Group("config").Use(middleware.JWTAuthFilter())
	{
		configRouter.POST("create", r.configApi.Create)        // 新建Config
		configRouter.DELETE("delete/:ids", r.configApi.Delete) // 删除Config
		configRouter.POST("update", r.configApi.Update)        // 更新Config
	}
	{
		configRouterWithoutRecord.GET("get/:id", r.configApi.Get)                      // 根据ID获取Config
		configRouterWithoutRecord.GET("configKey/:key", r.configApi.SelectConfigByKey) // 根据ID获取Config
		configRouterWithoutRecord.GET("page", r.configApi.Page)                        // 分页获取Config列表
		configRouterWithoutRecord.GET("list", r.configApi.List)                        // 分页获取Config列表
	}
}
