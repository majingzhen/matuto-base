// Package api  SysConfigApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_config/service"
	"manager-gin/src/app/admin/sys/sys_config/service/view"
	response "manager-gin/src/common/response"
	"manager-gin/src/framework"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strings"
)

type SysConfigApi struct {
}

var sysConfigService = service.SysConfigServiceApp

// Create 创建SysConfig
// @Summary 创建SysConfig
// @Router /sysConfig/create [post]
func (api *SysConfigApi) Create(c *gin.Context) {
	var sysConfigView view.SysConfigView
	_ = c.ShouldBindJSON(&sysConfigView)
	sysConfigView.Id = utils.GenUID()
	sysConfigView.CreateTime = utils.GetCurTimeStr()
	sysConfigView.UpdateTime = utils.GetCurTimeStr()
	sysConfigView.CreateBy = framework.GetLoginUserName(c)
	if err := sysConfigService.Create(&sysConfigView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除SysConfig
// @Summary 删除SysConfig
// @Router /sysConfig/delete [delete]
func (api *SysConfigApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := sysConfigService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新SysConfig
// @Summary 更新SysConfig
// @Router /sysConfig/update [put]
func (api *SysConfigApi) Update(c *gin.Context) {
	var sysConfigView view.SysConfigView
	_ = c.ShouldBindJSON(&sysConfigView)
	id := sysConfigView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	sysConfigView.UpdateTime = utils.GetCurTimeStr()
	sysConfigView.UpdateBy = framework.GetLoginUserName(c)
	if err := sysConfigService.Update(id, &sysConfigView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysConfig
// @Summary 用id查询SysConfig
// @Router /sysConfig/get [get]
func (api *SysConfigApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, sysConfigView := sysConfigService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysConfigView, c)
	}
}

// Page 分页获取SysConfig列表
// @Summary 分页获取SysConfig列表
// @Router /sysConfig/page [get]
func (api *SysConfigApi) Page(c *gin.Context) {
	var pageInfo view.SysConfigPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	if err, res := sysConfigService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取SysConfig列表
// @Summary 获取SysConfig列表
// @Router /sysConfig/list [get]
func (api *SysConfigApi) List(c *gin.Context) {
	var view view.SysConfigView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := sysConfigService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectConfigByKey 根据key查询SysMenu
// @Summary 根据key查询SysMenu
// @Router /sysConfig/selectConfigByKey [get]
func (api *SysConfigApi) SelectConfigByKey(c *gin.Context) {
	key := c.Param("key")
	if err, sysConfigView := sysConfigService.SelectConfigByKey(key); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(sysConfigView, c)
	}
}
