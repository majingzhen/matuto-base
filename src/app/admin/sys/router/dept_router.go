// Package router DeptRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type DeptRouter struct {
	deptApi api.DeptApi
}

// InitDeptRouter 初始化 Dept 路由信息
func (r *DeptRouter) InitDeptRouter(Router *gin.RouterGroup) {
	deptRouter := Router.Group("dept").Use(middleware.JWTAuthFilter())
	deptRouterWithoutRecord := Router.Group("dept").Use(middleware.JWTAuthFilter())
	{
		deptRouter.POST("create", r.deptApi.Create)       // 新建Dept
		deptRouter.DELETE("delete/:id", r.deptApi.Delete) // 删除Dept
		deptRouter.POST("update", r.deptApi.Update)       // 更新Dept
	}
	{
		deptRouterWithoutRecord.GET("get/:id", r.deptApi.Get) // 根据ID获取Dept
		deptRouterWithoutRecord.GET("list", r.deptApi.List)   // 获取Dept列表
		deptRouterWithoutRecord.GET("list/exclude/:id", r.deptApi.ListExclude)
		deptRouterWithoutRecord.GET("tree", r.deptApi.SelectDeptTree)                     // 获取部门树
		deptRouterWithoutRecord.GET("treeByRole/:roleId", r.deptApi.SelectDeptTreeByRole) // 根据角色ID获取部门树
	}
}
