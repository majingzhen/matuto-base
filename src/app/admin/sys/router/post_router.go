// Package router PostRouter 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package router

import (
	"github.com/gin-gonic/gin"
	"matuto-base/src/app/admin/sys/api"
	"matuto-base/src/middleware"
)

type PostRouter struct {
	postApi api.PostApi
}

// InitPostRouter 初始化 Post 路由信息
func (r *PostRouter) InitPostRouter(Router *gin.RouterGroup) {
	postRouter := Router.Group("post").Use(middleware.JWTAuthFilter())
	postRouterWithoutRecord := Router.Group("post").Use(middleware.JWTAuthFilter())
	{
		postRouter.POST("create", r.postApi.Create)        // 新建Post
		postRouter.DELETE("delete/:ids", r.postApi.Delete) // 删除Post
		postRouter.POST("update", r.postApi.Update)        // 更新Post
	}
	{
		postRouterWithoutRecord.GET("get/:id", r.postApi.Get) // 根据ID获取Post
		postRouterWithoutRecord.GET("page", r.postApi.Page)   // 分页获取Post列表
		postRouterWithoutRecord.GET("list", r.postApi.List)   // 分页获取Post列表
	}
}
