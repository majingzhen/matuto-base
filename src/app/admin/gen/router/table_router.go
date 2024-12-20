// Package router TableRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: gen_table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/gen/api"
	"matuto-base/src/middleware"
)

type TableRouter struct {
	tableApi api.TableApi
}

// InitTableRouter 初始化 Table 路由信息
func (r *TableRouter) InitTableRouter(Router *gin.RouterGroup) {
	tableRouter := Router.Group("table").Use(middleware.JWTAuthFilter())
	tableRouterWithoutRecord := Router.Group("table").Use(middleware.JWTAuthFilter())
	{
		tableRouter.POST("create", r.tableApi.Create)           // 新建Table
		tableRouter.DELETE("delete/:ids", r.tableApi.Delete)    // 删除Table
		tableRouter.PUT("update", r.tableApi.Update)            // 更新Table
		tableRouter.POST("importTable", r.tableApi.ImportTable) // 导入table
	}
	{
		tableRouterWithoutRecord.GET("get/:id", r.tableApi.Get)               // 根据ID获取Table
		tableRouterWithoutRecord.GET("page", r.tableApi.Page)                 // 分页获取Table列表
		tableRouterWithoutRecord.GET("db/page", r.tableApi.SelectDbTablePage) // 分页获取未导入的Table列表
		tableRouterWithoutRecord.GET("list", r.tableApi.List)                 // 分页获取Table列表
		tableRouterWithoutRecord.GET("preview/:id", r.tableApi.Preview)       // 预览代码
		tableRouterWithoutRecord.GET("genCode/:id", r.tableApi.GenCode)       // 生成代码
	}
}
