// Package api  DictTypeApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/service/dict_type"
	"matuto-base/src/app/admin/sys/service/dict_type/view"
	"matuto-base/src/common/basic"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type DictTypeApi struct {
	basic.BasicApi
	dictTypeService dict_type.DictTypeService
}

// Create 创建DictType
// @Summary 创建DictType
// @Router /dictType/create [post]
func (api *DictTypeApi) Create(c *gin.Context) {
	var dictTypeView view.DictTypeView
	_ = c.ShouldBindJSON(&dictTypeView)
	dictTypeView.Id = utils.GenUID()
	dictTypeView.CreateTime = utils.GetCurTimeStr()
	dictTypeView.UpdateTime = utils.GetCurTimeStr()
	dictTypeView.CreateBy = api.GetLoginUserName(c)
	if err := api.dictTypeService.Create(&dictTypeView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除DictType
// @Summary 删除DictType
// @Router /dictType/delete [delete]
func (api *DictTypeApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := api.dictTypeService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新DictType
// @Summary 更新DictType
// @Router /dictType/update [put]
func (api *DictTypeApi) Update(c *gin.Context) {
	var dictTypeView view.DictTypeView
	_ = c.ShouldBindJSON(&dictTypeView)
	id := dictTypeView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	dictTypeView.UpdateTime = utils.GetCurTimeStr()
	dictTypeView.UpdateBy = api.GetLoginUserName(c)
	if err := api.dictTypeService.Update(id, &dictTypeView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询DictType
// @Summary 用id查询DictType
// @Router /dictType/get [get]
func (api *DictTypeApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, dictTypeView := api.dictTypeService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dictTypeView, c)
	}
}

// Page 分页获取DictType列表
// @Summary 分页获取DictType列表
// @Router /dictType/list [get]
func (api *DictTypeApi) Page(c *gin.Context) {
	var pageInfo view.DictTypePageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	if err, res := api.dictTypeService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectDictTypeAll 获取DictType全部数据
// @Summary 获取DictType全部数据
// @Router /dictType/list [get]
func (api *DictTypeApi) SelectDictTypeAll(c *gin.Context) {
	if err, res := api.dictTypeService.SelectDictTypeAll(); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
