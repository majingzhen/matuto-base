// Package api  PostApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/service/post"
	"matuto-base/src/common/basic"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

// PostApi 结构体
type PostApi struct {
	basic.BasicApi
	postService post.PostService
}

// Create 创建Post
// @Summary 创建Post
// @Router /post/create [post]
func (api *PostApi) Create(c *gin.Context) {
	var postView vo.PostView
	_ = c.ShouldBindJSON(&postView)
	postView.Id = utils.GenUID()
	postView.CreateTime = utils.GetCurTimeStr()
	postView.UpdateTime = utils.GetCurTimeStr()
	postView.CreateBy = api.GetLoginUserName(c)
	if err := api.postService.Create(&postView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除Post
// @Summary 删除Post
// @Router /post/delete [delete]
func (api *PostApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := api.postService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新Post
// @Summary 更新Post
// @Router /post/update [put]
func (api *PostApi) Update(c *gin.Context) {
	var postView vo.PostView
	_ = c.ShouldBindJSON(&postView)
	id := postView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	postView.UpdateTime = utils.GetCurTimeStr()
	postView.UpdateBy = api.GetLoginUserName(c)
	if err := api.postService.Update(id, &postView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询Post
// @Summary 用id查询Post
// @Router /post/get [get]
func (api *PostApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, postView := api.postService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(postView, c)
	}
}

// Page 分页获取Post列表
// @Summary 分页获取Post列表
// @Router /post/page [get]
func (api *PostApi) Page(c *gin.Context) {
	var pageInfo vo.PostPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}

	if err, res := api.postService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取Post列表
// @Summary 获取Post列表
// @Router /post/list [get]
func (api *PostApi) List(c *gin.Context) {
	var view vo.PostView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := api.postService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
