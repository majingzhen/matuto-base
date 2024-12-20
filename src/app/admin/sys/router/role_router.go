// Package router RoleRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type RoleRouter struct {
	roleApi api.RoleApi
}

// InitRoleRouter 初始化 Role 路由信息
func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role").Use(middleware.JWTAuthFilter())
	roleRouterWithoutRecord := Router.Group("role").Use(middleware.JWTAuthFilter())
	{
		roleRouter.POST("create", r.roleApi.Create)                      // 新建Role
		roleRouter.DELETE("delete/:ids", r.roleApi.Delete)               // 删除Role
		roleRouter.POST("update", r.roleApi.Update)                      // 更新Role
		roleRouter.PUT("changeStatus", r.roleApi.ChangeStatus)           // 更新Role状态
		roleRouter.PUT("dataScope", r.roleApi.DataScope)                 // 数据授权
		roleRouter.PUT("cancelAuth", r.roleApi.CancelAuthUser)           // 取消授权
		roleRouter.PUT("batchCancelAuth", r.roleApi.BatchCancelAuthUser) // 批量取消授权
		roleRouter.PUT("batchSelectAuth", r.roleApi.BatchSelectAuthUser) // 批量选择授权

	}
	{
		roleRouterWithoutRecord.GET("get/:id", r.roleApi.Get) // 根据ID获取Role
		roleRouterWithoutRecord.GET("page", r.roleApi.Page)   // 分页获取Role列表
		roleRouterWithoutRecord.GET("list", r.roleApi.List)   // 分页获取Role列表
	}
}
