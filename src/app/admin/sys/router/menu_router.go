// Package router MenuRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type MenuRouter struct {
	menuApi api.MenuApi
}

// InitMenuRouter 初始化 Menu 路由信息
func (r *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu").Use(middleware.JWTAuthFilter())
	menuRouterWithoutRecord := Router.Group("menu").Use(middleware.JWTAuthFilter())
	{
		menuRouter.POST("create", r.menuApi.Create)       // 新建Menu
		menuRouter.DELETE("delete/:id", r.menuApi.Delete) // 删除Menu
		menuRouter.POST("update", r.menuApi.Update)       // 更新Menu
	}
	{
		menuRouterWithoutRecord.GET("get/:id", r.menuApi.Get)                                       // 根据ID获取Menu
		menuRouterWithoutRecord.GET("treeSelect", r.menuApi.SelectMenuTree)                         // 获取Menu树形列表
		menuRouterWithoutRecord.GET("roleMenuTreeSelect/:roleId", r.menuApi.SelectMenuTreeByRoleId) // 根据ID获取Menu
		menuRouterWithoutRecord.GET("list", r.menuApi.List)                                         // 获取Menu列表
	}
}
