// Package api  ConfigApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/service/config"
	"matuto-base/src/app/admin/sys/service/config/view"
	"matuto-base/src/common/basic"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type ConfigApi struct {
	basic.BasicApi
	configService config.ConfigService
}

// Create 创建Config
// @Summary 创建Config
// @Router /config/create [post]
func (api *ConfigApi) Create(c *gin.Context) {
	var configView view.ConfigView
	_ = c.ShouldBindJSON(&configView)
	// 判断是否重复
	if err, value := api.configService.SelectConfigByKey(configView.ConfigKey); err != nil {
		response.FailWithMessage("更新失败", c)
		return
	} else {
		if value != nil {
			response.FailWithMessage("配置键名已存在", c)
			return
		}
	}
	configView.Id = utils.GenUID()
	configView.CreateTime = utils.GetCurTimeStr()
	configView.UpdateTime = utils.GetCurTimeStr()
	configView.CreateBy = api.GetLoginUserName(c)
	if err := api.configService.Create(&configView); err != nil {
		global.Logger.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除Config
// @Summary 删除Config
// @Router /config/delete [delete]
func (api *ConfigApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := api.configService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新Config
// @Summary 更新Config
// @Router /config/update [put]
func (api *ConfigApi) Update(c *gin.Context) {
	var configView view.ConfigView
	_ = c.ShouldBindJSON(&configView)
	id := configView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 判断是否重复
	if err, value := api.configService.SelectConfigByKey(configView.ConfigKey); err != nil {
		response.FailWithMessage("更新失败", c)
		return
	} else {
		if value != nil && value.Id != configView.Id {
			response.FailWithMessage("配置键名已存在", c)
			return
		}
	}
	configView.UpdateTime = utils.GetCurTimeStr()
	configView.UpdateBy = api.GetLoginUserName(c)
	if err := api.configService.Update(id, &configView); err != nil {
		global.Logger.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询Config
// @Summary 用id查询Config
// @Router /config/get [get]
func (api *ConfigApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, configView := api.configService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(configView, c)
	}
}

// Page 分页获取Config列表
// @Summary 分页获取Config列表
// @Router /config/page [get]
func (api *ConfigApi) Page(c *gin.Context) {
	var pageInfo view.ConfigPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	if err, res := api.configService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取Config列表
// @Summary 获取Config列表
// @Router /config/list [get]
func (api *ConfigApi) List(c *gin.Context) {
	var view view.ConfigView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := framework.GetLoginUserId(c)
	if err, res := api.configService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectConfigByKey 根据key查询Menu
// @Summary 根据key查询Menu
// @Router /config/selectConfigByKey [get]
func (api *ConfigApi) SelectConfigByKey(c *gin.Context) {
	key := c.Param("key")
	if err, configView := api.configService.SelectConfigByKey(key); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(configView.ConfigValue, c)
	}
}
