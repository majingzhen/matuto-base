// Package api  MenuApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: menu
// @version 1.0.0
// @create 2023-08-20 21:21:34
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/service"
	"matuto-base/src/common/basic"
	"matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type MenuApi struct {
	basic.BasicApi
	menuService service.MenuService
}

// Create 创建Menu
// @Summary 创建Menu
// @Router /menu/create [post]
func (api *MenuApi) Create(c *gin.Context) {
	var menuView vo.MenuView
	_ = c.ShouldBindJSON(&menuView)
	menuView.Id = utils.GenUID()
	menuView.CreateTime = utils.GetCurTimeStr()
	menuView.UpdateTime = utils.GetCurTimeStr()
	menuView.CreateBy = api.GetLoginUserName(c)
	if err := api.menuService.Create(&menuView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除Menu
// @Summary 删除Menu
// @Router /menu/delete [delete]
func (api *MenuApi) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := api.menuService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新Menu
// @Summary 更新Menu
// @Router /menu/update [put]
func (api *MenuApi) Update(c *gin.Context) {
	var menuView vo.MenuView
	_ = c.ShouldBindJSON(&menuView)
	id := menuView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	menuView.UpdateTime = utils.GetCurTimeStr()
	menuView.UpdateBy = api.GetLoginUserName(c)
	if err := api.menuService.Update(id, &menuView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询Menu
// @Summary 用id查询Menu
// @Router /menu/get [get]
func (api *MenuApi) Get(c *gin.Context) {
	id := c.Param("id")
	if err, menuView := api.menuService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(menuView, c)
	}
}

// List 获取Menu列表
// @Summary 获取Menu列表
// @Router /menu/list [get]
func (api *MenuApi) List(c *gin.Context) {
	var menuView vo.MenuView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&menuView); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	userId := api.GetLoginUserId(c)
	if err, res := api.menuService.SelectMenuList(&menuView, userId); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectMenuTreeByRoleId 加载对应角色菜单列表树
// @Summary 根据角色id查询菜单树信息
// @Router /menu/selectMenuTreeByRoleId/{roleId} [get]
func (api *MenuApi) SelectMenuTreeByRoleId(c *gin.Context) {
	roleId := c.Param("roleId")
	if err, menuList := api.menuService.SelectMenuList(&vo.MenuView{}, api.GetLoginUserId(c)); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		menuTree := api.menuService.BuildMenuTreeSelect(menuList)
		if err, menuListByRoleId := api.menuService.SelectMenuListByRoleId(roleId); err != nil {
			global.Logger.Error("获取数据失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
		} else {
			response.OkWithData(gin.H{
				"checkedKeys": menuListByRoleId,
				"menus":       menuTree,
			}, c)
		}
	}
}

// SelectMenuTree 加载菜单列表树
// @Summary 加载菜单列表树
// @Router /menu/selectMenuTree [get]
func (api *MenuApi) SelectMenuTree(c *gin.Context) {
	var menuView vo.MenuView
	if err := c.ShouldBindQuery(&menuView); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	if err, menuList := api.menuService.SelectMenuList(&menuView, api.GetLoginUserId(c)); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		menuTree := api.menuService.BuildMenuTreeSelect(menuList)
		response.OkWithData(menuTree, c)
	}
}
