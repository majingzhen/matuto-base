// Package api  NoticeApi 自动生成模板
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-12 13:58:38
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/service/notice"
	"matuto-base/src/app/admin/sys/service/notice/view"
	"matuto-base/src/common/basic"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type NoticeApi struct {
	basic.BasicApi
	noticeService notice.NoticeService
}

// Create 创建Notice
// @Summary 创建Notice
// @Router /notice/create [post]
func (api *NoticeApi) Create(c *gin.Context) {
	var noticeView view.NoticeCreateView
	_ = c.ShouldBindJSON(&noticeView)
	noticeView.Id = utils.GenUID()
	noticeView.CreateTime = utils.GetCurTimeStr()
	noticeView.UpdateTime = utils.GetCurTimeStr()
	noticeView.CreateBy = api.GetLoginUserName(c)
	if err := api.noticeService.Create(&noticeView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除Notice
// @Summary 删除Notice
// @Router /notice/delete [delete]
func (api *NoticeApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	if idStr == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	ids := strings.Split(idStr, ",")
	if err := api.noticeService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新Notice
// @Summary 更新Notice
// @Router /notice/update [put]
func (api *NoticeApi) Update(c *gin.Context) {
	var noticeView view.NoticeEditView
	err := c.ShouldBindJSON(&noticeView)
	if err != nil {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	noticeView.UpdateTime = utils.GetCurTimeStr()
	noticeView.UpdateBy = api.GetLoginUserName(c)
	if err := api.noticeService.Update(&noticeView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询Notice
// @Summary 用id查询Notice
// @Router /notice/get [get]
func (api *NoticeApi) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	if err, noticeView := api.noticeService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(noticeView, c)
	}
}

// Page 分页获取Notice列表
// @Summary 分页获取Notice列表
// @Router /notice/page [get]
func (api *NoticeApi) Page(c *gin.Context) {
	var pageInfo view.NoticePageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("参数解析失败!", c)
		return
	}
	if err, res := api.noticeService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取Notice列表
// @Summary 获取Notice列表
// @Router /notice/list [get]
func (api *NoticeApi) List(c *gin.Context) {
	var view view.NoticeQueryView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := api.noticeService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
