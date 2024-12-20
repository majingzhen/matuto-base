// Package router DictTypeRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type DictTypeRouter struct {
	dictTypeApi api.DictTypeApi
}

// InitDictTypeRouter 初始化 DictType 路由信息
func (r *DictTypeRouter) InitDictTypeRouter(Router *gin.RouterGroup) {
	dictTypeRouter := Router.Group("dictType").Use(middleware.JWTAuthFilter())
	dictTypeRouterWithoutRecord := Router.Group("dictType").Use(middleware.JWTAuthFilter())
	{
		dictTypeRouter.POST("create", r.dictTypeApi.Create)        // 新建DictType
		dictTypeRouter.DELETE("delete/:ids", r.dictTypeApi.Delete) // 删除DictType
		dictTypeRouter.POST("update", r.dictTypeApi.Update)        // 更新DictType
	}
	{
		dictTypeRouterWithoutRecord.GET("get/:id", r.dictTypeApi.Get)                    // 根据ID获取DictType
		dictTypeRouterWithoutRecord.GET("page", r.dictTypeApi.Page)                      // 获取DictType列表
		dictTypeRouterWithoutRecord.GET("optionSelect", r.dictTypeApi.SelectDictTypeAll) // 获取DictType列表
	}
}
