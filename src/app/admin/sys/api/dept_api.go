// Package api  DeptApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/service/dept"
	"matuto-base/src/app/admin/sys/service/dept/view"
	"matuto-base/src/common/basic"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type DeptApi struct {
	basic.BasicApi
	deptService dept.DeptService
}

// Create 创建Dept
// @Summary 创建Dept
// @Router /dept/create [post]
func (api *DeptApi) Create(c *gin.Context) {
	var deptView view.DeptView
	_ = c.ShouldBindJSON(&deptView)
	deptView.Id = utils.GenUID()
	deptView.CreateTime = utils.GetCurTimeStr()
	deptView.UpdateTime = utils.GetCurTimeStr()
	deptView.CreateBy = api.GetLoginUserName(c)
	if err := api.deptService.Create(&deptView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除Dept
// @Summary 删除Dept
// @Router /dept/delete [delete]
func (api *DeptApi) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := api.deptService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新Dept
// @Summary 更新Dept
// @Router /dept/update [put]
func (api *DeptApi) Update(c *gin.Context) {
	var deptView view.DeptView
	_ = c.ShouldBindJSON(&deptView)
	id := deptView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	deptView.UpdateTime = utils.GetCurTimeStr()
	deptView.UpdateBy = api.GetLoginUser(c).UserName
	if err := api.deptService.Update(id, &deptView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询Dept
// @Summary 用id查询Dept
// @Router /dept/get [get]
func (api *DeptApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, deptView := api.deptService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(deptView, c)
	}
}

// List 获取Dept列表
// @Summary 获取Dept列表
// @Router /dept/list [get]
func (api *DeptApi) List(c *gin.Context) {
	var view view.DeptView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := api.GetLoginUserId(c)
	user := api.GetLoginUser(c)
	if err, res := api.deptService.List(&view, user); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// ListExclude 查询部门列表（排除节点）
// @Summary 查询部门列表（排除节点）
// @Router /dept/listExclude [get]
func (api *DeptApi) ListExclude(c *gin.Context) {
	id := c.Param("id")
	user := api.GetLoginUser(c)
	if err, deptView := api.deptService.List(&view.DeptView{}, user); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		var deleteIndexes []int // 用于记录要删除的元素索引
		for i, deptView := range deptView {
			parentIds := strings.Split(deptView.Ancestors, ",")
			// 要把查询的这个节点排除掉
			if deptView.Id == id || utils.ContainsStr(parentIds, id) {
				deleteIndexes = append(deleteIndexes, i)
			}
		}
		// 根据记录的索引删除元素
		for i := len(deleteIndexes) - 1; i >= 0; i-- {
			deptView = append(deptView[:deleteIndexes[i]], deptView[deleteIndexes[i]+1:]...)
		}
		response.OkWithData(deptView, c)
	}
}

// SelectDeptTree 查询部门树列表
// @Summary 查询部门树列表
// @Router /dept/selectDeptTree [get]
func (api *DeptApi) SelectDeptTree(c *gin.Context) {
	var view view.DeptView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, deptView := api.deptService.SelectDeptTree(&view, user); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(deptView, c)
	}
}

// SelectDeptTreeByRole 查询部门树列表
// @Summary 根据角色查询部门树列表
// @Router /dept/selectDeptTreeByRole [get]
func (api *DeptApi) SelectDeptTreeByRole(c *gin.Context) {
	roleId := c.Param("roleId")
	if roleId == "" {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	if err, checkedKeys := api.deptService.SelectDeptTreeByRole(roleId); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		err1, deptTree := api.deptService.SelectDeptTree(&view.DeptView{}, api.GetLoginUser(c))
		if err1 != nil {
			global.Logger.Error("查询失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(gin.H{
			"checkedKeys": checkedKeys,
			"depts":       deptTree,
		}, c)
	}
}
