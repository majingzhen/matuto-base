// Package router TableColumnRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/gen/api"
	"matuto-base/src/middleware"
)

type TableColumnRouter struct {
	tableColumnApi api.TableColumnApi
}

// InitTableColumnRouter 初始化 TableColumn 路由信息
func (r *TableColumnRouter) InitTableColumnRouter(Router *gin.RouterGroup) {
	tableColumnRouter := Router.Group("tableColumn").Use(middleware.JWTAuthFilter())
	tableColumnRouterWithoutRecord := Router.Group("tableColumn").Use(middleware.JWTAuthFilter())
	{
		tableColumnRouter.DELETE("delete/:ids", r.tableColumnApi.Delete) // 删除TableColumn
		tableColumnRouter.PUT("update", r.tableColumnApi.Update)         // 更新TableColumn
	}
	{
		tableColumnRouterWithoutRecord.GET("get/:id", r.tableColumnApi.Get) // 根据ID获取TableColumn
	}
}
