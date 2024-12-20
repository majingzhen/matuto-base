// Package api  RoleApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/service/role"
	"matuto-base/src/common/basic"
	"matuto-base/src/common/constants"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type RoleApi struct {
	basic.BasicApi
	roleService role.RoleService
}

// Create 创建Role
// @Summary 创建Role
// @Router /role/create [post]
func (api *RoleApi) Create(c *gin.Context) {
	var roleView vo.RoleView
	if err := c.ShouldBindJSON(&roleView); err != nil {
		global.Logger.Error("参数解析失败!", zap.Error(err))
		response.FailWithMessage("参数解析失败", c)
		return
	}
	if err := api.roleService.CheckRoleNameUnique(roleView.RoleName, ""); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := api.roleService.CheckRoleKeyUnique(roleView.RoleKey, ""); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	roleView.Id = utils.GenUID()
	roleView.CreateTime = utils.GetCurTimeStr()
	roleView.UpdateTime = utils.GetCurTimeStr()
	roleView.CreateBy = api.GetLoginUserName(c)
	if err := api.roleService.Create(&roleView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除Role
// @Summary 删除Role
// @Router /role/delete [delete]
func (api *RoleApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if err := api.roleService.DeleteByIds(ids, api.GetLoginUser(c)); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新Role
// @Summary 更新Role
// @Router /role/update [put]
func (api *RoleApi) Update(c *gin.Context) {
	var roleView vo.RoleView
	_ = c.ShouldBindJSON(&roleView)
	id := roleView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 校验参数
	if id == constants.SYSTEM_ROLE_ADMIN_ID {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	// 校验数据权限
	if err := api.roleService.CheckRoleDataScope(id, api.GetLoginUser(c)); err != nil {
		response.FailWithMessage("没有权限访问角色数据", c)
		return
	}
	if err := api.roleService.CheckRoleNameUnique(roleView.RoleName, roleView.Id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := api.roleService.CheckRoleKeyUnique(roleView.RoleKey, roleView.Id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	roleView.UpdateTime = utils.GetCurTimeStr()
	roleView.UpdateBy = api.GetLoginUserName(c)
	if err := api.roleService.Update(id, &roleView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询Role
// @Summary 用id查询Role
// @Router /role/get [get]
func (api *RoleApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, roleView := api.roleService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(roleView, c)
	}
}

// Page 分页获取Role列表
// @Summary 分页获取Role列表
// @Router /role/page [get]
func (api *RoleApi) Page(c *gin.Context) {
	var pageInfo vo.RolePageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, res := api.roleService.Page(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取Role列表
// @Summary 获取Role列表
// @Router /role/list [get]
func (api *RoleApi) List(c *gin.Context) {
	var view vo.RoleView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := api.GetLoginUserId(c)
	if err, res := api.roleService.List(&view, api.GetLoginUser(c)); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// ChangeStatus 更新Role状态
// @Summary 更新Role状态
// @Router /role/changeStatus [put]
func (api *RoleApi) ChangeStatus(c *gin.Context) {
	var view vo.RoleView
	_ = c.ShouldBindJSON(&view)
	if view.Id == constants.SYSTEM_ROLE_ADMIN_ID {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.roleService.CheckRoleDataScope(view.Id, api.GetLoginUser(c)); err != nil {
		response.FailWithMessage("没有权限访问角色数据", c)
		return
	}
	view.UpdateTime = utils.GetCurTimeStr()
	view.UpdateBy = api.GetLoginUserName(c)
	if err := api.roleService.UpdateStatus(&view); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// DataScope 修改保存数据权限
// @Summary 修改保存数据权限
// @Router /role/dataScope [put]
func (api *RoleApi) DataScope(c *gin.Context) {
	var view vo.RoleView
	_ = c.ShouldBindJSON(&view)
	if view.Id == constants.SYSTEM_ROLE_ADMIN_ID {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.roleService.CheckRoleDataScope(view.Id, api.GetLoginUser(c)); err != nil {
		response.FailWithMessage("没有权限访问角色数据", c)
		return
	}
	view.UpdateTime = utils.GetCurTimeStr()
	view.UpdateBy = api.GetLoginUserName(c)
	if err := api.roleService.AuthDataScope(&view); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CancelAuthUser 取消授权用户
// @Summary 取消授权用户
// @Router /role/cancelAuthUser [put]
func (api *RoleApi) CancelAuthUser(c *gin.Context) {
	var view vo.UserRoleView
	_ = c.ShouldBindJSON(&view)
	if err := api.roleService.CancelAuthUser(&view); err != nil {
		global.Logger.Error("取消授权失败!", zap.Error(err))
		response.FailWithMessage("取消授权失败", c)
	} else {
		response.OkWithMessage("取消授权成功", c)
	}
}

// BatchCancelAuthUser 批量取消授权用户
// @Summary 批量取消授权用户
// @Router /role/batchCancelAuthUser [put]
func (api *RoleApi) BatchCancelAuthUser(c *gin.Context) {
	userId := c.Query("userIds")
	roleId := c.Query("roleId")
	if userId == "" || roleId == "" {
		response.FailWithMessage("参数解析错误", c)
	}
	userIds := strings.Split(userId, ",")
	if err := api.roleService.BatchCancelAuthUser(roleId, userIds); err != nil {
		global.Logger.Error("批量取消授权失败!", zap.Error(err))
		response.FailWithMessage("批量取消授权失败", c)
	} else {
		response.OkWithMessage("批量取消授权成功", c)
	}
}

// BatchSelectAuthUser 批量授权用户
// @Summary 授权用户
// @Router /role/batchSelectAuthUser [put]
func (api *RoleApi) BatchSelectAuthUser(c *gin.Context) {
	userId := c.Query("userIds")
	roleId := c.Query("roleId")
	if userId == "" || roleId == "" {
		response.FailWithMessage("参数解析错误", c)
	}
	userIds := strings.Split(userId, ",")
	if err := api.roleService.BatchSelectAuthUser(roleId, userIds); err != nil {
		global.Logger.Error("批量授权失败!", zap.Error(err))
		response.FailWithMessage("批量授权失败", c)
	} else {
		response.OkWithMessage("批量授权成功", c)
	}
}
