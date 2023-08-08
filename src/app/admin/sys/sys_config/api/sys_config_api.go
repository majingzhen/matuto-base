// Package api  SysConfigApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_config
// @version 1.0.0
// @create 2023-08-08 10:06:19
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_config/service"
	"manager-gin/src/app/admin/sys/sys_config/service/view"
	"manager-gin/src/common"
	"manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"strconv"
)

type SysConfigApi struct {
}

var sysConfigService = service.SysConfigServiceApp

// Create 创建SysConfig
// @Summary 创建SysConfig
// @Router /sysConfig/create [post]
func (sysConfigApi *SysConfigApi) Create(c *gin.Context) {
	var sysConfigView view.SysConfigView
	_ = c.ShouldBindJSON(&sysConfigView)
	sysConfigView.CreateTime = utils.GetCurTimeStr()
	sysConfigView.UpdateTime = utils.GetCurTimeStr()
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
func (sysConfigApi *SysConfigApi) Delete(c *gin.Context) {
	var id int
	_ = c.ShouldBindJSON(&id)
	if err := sysConfigService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除SysConfig
// @Summary 批量删除SysConfig
// @Router /sysConfig/deleteByIds [delete]
func (sysConfigApi *SysConfigApi) DeleteByIds(c *gin.Context) {
	var ids []int
	_ = c.ShouldBindJSON(&ids)
	if err := sysConfigService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新SysConfig
// @Summary 更新SysConfig
// @Router /sysConfig/update [put]
func (sysConfigApi *SysConfigApi) Update(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	sysConfigViewJson := c.Query("sysConfigView")
	var sysConfigView view.SysConfigView
	err = json.Unmarshal([]byte(sysConfigViewJson), &sysConfigView)
	sysConfigView.UpdateTime = utils.GetCurTimeStr()
	if err := sysConfigService.Update(atoi, &sysConfigView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = sysConfigService.Update(atoi, &sysConfigView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询SysConfig
// @Summary 用id查询SysConfig
// @Router /sysConfig/get [get]
func (sysConfigApi *SysConfigApi) Get(c *gin.Context) {
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, sysConfigView := sysConfigService.Get(atoi); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"sysConfigView": sysConfigView}, c)
	}
}

// Find 分页获取SysConfig列表
// @Summary 分页获取SysConfig列表
// @Router /sysConfig/find [get]
func (sysConfigApi *SysConfigApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	_ = c.ShouldBindQuery(&pageInfo)
	if err := sysConfigService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
