// Package router DictDataRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type DictDataRouter struct {
	dictDataApi api.DictDataApi
}

// InitDictDataRouter 初始化 DictData 路由信息
func (r *DictDataRouter) InitDictDataRouter(Router *gin.RouterGroup) {
	dictDataRouter := Router.Group("dictData").Use(middleware.JWTAuthFilter())
	dictDataRouterWithoutRecord := Router.Group("dictData").Use(middleware.JWTAuthFilter())
	{
		dictDataRouter.POST("create", r.dictDataApi.Create)        // 新建DictData
		dictDataRouter.DELETE("delete/:ids", r.dictDataApi.Delete) // 删除DictData
		dictDataRouter.POST("update", r.dictDataApi.Update)        // 更新DictData
	}
	{
		dictDataRouterWithoutRecord.GET("get/:id", r.dictDataApi.Get)          // 根据ID获取DictData
		dictDataRouterWithoutRecord.GET("type/:type", r.dictDataApi.GetByType) // 根据type获取DictData
		dictDataRouterWithoutRecord.GET("page", r.dictDataApi.Page)            // 获取DictData列表
	}
}
