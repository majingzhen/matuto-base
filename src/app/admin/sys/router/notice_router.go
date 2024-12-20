// Package router NoticeRouter 自动生成模板
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-12 13:58:38
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

// NoticeRouter 结构体
type NoticeRouter struct {
	noticeApi api.NoticeApi
}

// InitNoticeRouter 初始化 Notice 路由信息
func (r *NoticeRouter) InitNoticeRouter(Router *gin.RouterGroup) {
	noticeRouter := Router.Group("notice").Use(middleware.JWTAuthFilter())
	noticeRouterWithoutRecord := Router.Group("notice").Use(middleware.JWTAuthFilter())
	{
		noticeRouter.POST("create", r.noticeApi.Create)        // 新建Notice
		noticeRouter.DELETE("delete/:ids", r.noticeApi.Delete) // 删除Notice
		noticeRouter.PUT("update", r.noticeApi.Update)         // 更新Notice
	}
	{
		noticeRouterWithoutRecord.GET("get/:id", r.noticeApi.Get) // 根据ID获取Notice
		noticeRouterWithoutRecord.GET("page", r.noticeApi.Page)   // 分页获取Notice列表
		noticeRouterWithoutRecord.GET("list", r.noticeApi.List)   // 分页获取Notice列表
	}
}
