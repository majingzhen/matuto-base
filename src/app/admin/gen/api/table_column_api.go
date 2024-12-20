// Package api  TableColumnApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: gen_table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/gen/api/vo"
	"matuto-base/src/app/admin/gen/service"
	"matuto-base/src/common/basic"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type TableColumnApi struct {
	basic.BasicApi
	tableColumnService service.TableColumnService
}

// Delete 删除TableColumn
// @Summary 删除TableColumn
// @Router /tableColumn/delete [delete]
func (api *TableColumnApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	if idStr == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	ids := strings.Split(idStr, ",")
	if err := api.tableColumnService.DeleteByIds(ids); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新TableColumn
// @Summary 更新TableColumn
// @Router /tableColumn/update [put]
func (api *TableColumnApi) Update(c *gin.Context) {
	var tableColumnView vo.TableColumnView
	err := c.ShouldBindJSON(&tableColumnView)
	if err != nil {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	tableColumnView.UpdateTime = utils.GetCurTimeStr()
	tableColumnView.UpdateBy = api.GetLoginUserName(c)
	if err := api.tableColumnService.Update(&tableColumnView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询TableColumn
// @Summary 用id查询TableColumn
// @Router /tableColumn/get [get]
func (api *TableColumnApi) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数解析错误", c)
		return
	}
	if err, tableColumnView := api.tableColumnService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(tableColumnView, c)
	}
}
