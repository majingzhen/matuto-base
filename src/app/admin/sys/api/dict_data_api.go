// Package api  DictDataApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/service/dict_data"
	"matuto-base/src/app/admin/sys/service/dict_data/view"
	"matuto-base/src/common/basic"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type DictDataApi struct {
	basic.BasicApi
	dictDataService dict_data.DictDataService
}

// Create 创建DictData
// @Summary 创建DictData
// @Router /dictData/create [post]
func (api *DictDataApi) Create(c *gin.Context) {
	var dictDataView view.DictDataView
	_ = c.ShouldBindJSON(&dictDataView)
	dictDataView.Id = utils.GenUID()
	dictDataView.CreateTime = utils.GetCurTimeStr()
	dictDataView.UpdateTime = utils.GetCurTimeStr()
	dictDataView.CreateBy = api.GetLoginUserName(c)
	if err := api.dictDataService.Create(&dictDataView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除DictData
// @Summary 删除DictData
// @Router /dictData/delete [delete]
func (api *DictDataApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := api.dictDataService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新DictData
// @Summary 更新DictData
// @Router /dictData/update [put]
func (api *DictDataApi) Update(c *gin.Context) {
	var dictDataView view.DictDataView
	_ = c.ShouldBindJSON(&dictDataView)
	id := dictDataView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	dictDataView.UpdateTime = utils.GetCurTimeStr()
	dictDataView.UpdateBy = api.GetLoginUserName(c)
	if err := api.dictDataService.Update(id, &dictDataView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询DictData
// @Summary 用id查询DictData
// @Router /dictData/get [get]
func (api *DictDataApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, dictDataView := api.dictDataService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dictDataView, c)
	}
}

// GetByType 用字典类型查询DictData
// @Summary 用字典类型查询DictData
// @Router /dictData/type [get]
func (api *DictDataApi) GetByType(c *gin.Context) {
	dictType := c.Param("type")
	if err, dictDataView := api.dictDataService.GetByType(dictType); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dictDataView, c)
	}
}

// Page 分页获取DictData列表
// @Summary 分页获取DictData列表
// @Router /dictData/list [get]
func (api *DictDataApi) Page(c *gin.Context) {
	var pageInfo view.DictDataPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	if err, res := api.dictDataService.Page(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
