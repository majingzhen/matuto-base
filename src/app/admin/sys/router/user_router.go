// Package router UserRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type UserRouter struct {
	userApi api.UserApi
}

// InitUserRouter 初始化 User 路由信息
func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.JWTAuthFilter())
	userRouterWithoutRecord := Router.Group("user").Use(middleware.JWTAuthFilter())
	{
		userRouter.POST("create", r.userApi.Create)            // 新建User
		userRouter.DELETE("delete/:ids", r.userApi.Delete)     // 删除User
		userRouter.PUT("update", r.userApi.Update)             // 更新User
		userRouter.PUT("changeStatus", r.userApi.ChangeStatus) // 修改状态
		userRouter.PUT("resetPwd", r.userApi.ResetPwd)         // 重置密码
		userRouter.PUT("authRole", r.userApi.AuthRole)         // 重置密码
	}
	{
		userRouterWithoutRecord.GET("get", r.userApi.Get)
		userRouterWithoutRecord.GET("get/:id", r.userApi.Get)                           // 根据ID获取User
		userRouterWithoutRecord.GET("authRole/:id", r.userApi.GetAuthRole)              // 根据ID获取授权集合
		userRouterWithoutRecord.GET("page", r.userApi.Page)                             // 分页获取User列表
		userRouterWithoutRecord.GET("list", r.userApi.List)                             // 分页获取User列表
		userRouterWithoutRecord.GET("allocatedList", r.userApi.SelectAllocatedList)     // 分页获取角色已授权用户列表
		userRouterWithoutRecord.GET("unallocatedList", r.userApi.SelectUnallocatedList) // 分页获取角色未授权用户列表
	}
}
