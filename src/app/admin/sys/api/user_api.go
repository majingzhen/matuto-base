// Package api  UserApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-21 14:20:37
package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/service"
	"matuto-base/src/common/basic"
	"matuto-base/src/common/constants"
	response "matuto-base/src/common/response"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
)

type UserApi struct {
	basic.BasicApi
	userService service.Service
	roleService service.RoleService
	postService service.PostService
}

// Create 创建User
// @Summary 创建User
// @Router /user/create [post]
func (api *UserApi) Create(c *gin.Context) {
	var userView vo.UserView
	_ = c.ShouldBindJSON(&userView)
	// 校验参数

	if err := api.userService.CheckFieldUnique("user_name", userView.UserName, ""); err != nil {
		response.FailWithMessage("登录账号已存在", c)
		return
	}
	if err := api.userService.CheckFieldUnique("phone_number", userView.PhoneNumber, ""); err != nil {
		response.FailWithMessage("手机号码已存在", c)
		return
	}
	if err := api.userService.CheckFieldUnique("email", userView.Email, ""); err != nil {
		response.FailWithMessage("邮箱账号已存在", c)
		return
	}
	userView.Id = utils.GenUID()
	userView.CreateTime = utils.GetCurTimeStr()
	userView.UpdateTime = utils.GetCurTimeStr()
	userView.Salt = utils.GenUID()
	userView.Password = utils.EncryptionPassword(userView.Password, userView.Salt)
	userView.CreateBy = api.GetLoginUserName(c)
	if err := api.userService.Create(&userView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除User
// @Summary 删除User
// @Router /user/delete [delete]
func (api *UserApi) Delete(c *gin.Context) {
	idStr := c.Param("ids")
	ids := strings.Split(idStr, ",")
	if utils.ContainsStr(ids, api.GetLoginUserId(c)) {
		response.FailWithMessage("当前用户不能删除", c)
	}
	if err := api.userService.DeleteByIds(ids, api.GetLoginUserId(c)); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 更新User
// @Summary 更新User
// @Router /user/update [put]
func (api *UserApi) Update(c *gin.Context) {
	var userView vo.UserView
	_ = c.ShouldBindJSON(&userView)
	id := userView.Id
	if id == "" {
		response.FailWithMessage("更新失败", c)
		return
	}
	if constants.SYSTEM_ADMIN_ID == id {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	// 校验用户是否有数据权限
	err := api.userService.CheckUserDataScope(id, api.GetLoginUserId(c))
	if err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	// 校验参数
	if err := api.userService.CheckFieldUnique("user_name", userView.UserName, userView.Id); err != nil {
		response.FailWithMessage("登录账号已存在", c)
		return
	}
	if err := api.userService.CheckFieldUnique("phone_number", userView.PhoneNumber, userView.Id); err != nil {
		response.FailWithMessage("手机号码已存在", c)
		return
	}
	if err := api.userService.CheckFieldUnique("email", userView.Email, userView.Id); err != nil {
		response.FailWithMessage("邮箱账号已存在", c)
		return
	}
	userView.UpdateTime = utils.GetCurTimeStr()
	userView.UpdateBy = api.GetLoginUserName(c)
	if err := api.userService.Update(id, &userView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询User
// @Summary 用id查询User
// @Router /user/get [get]
func (api *UserApi) Get(c *gin.Context) {
	var userInfoView = new(vo.UserInfoView)
	err, roles := api.roleService.SelectRoleAll(api.GetLoginUser(c))
	if err == nil {
		removeAdminRole(&roles)
		userInfoView.Roles = roles
	}
	err, views := api.postService.SelectPostAll()
	if err == nil {
		userInfoView.Posts = views
	}
	id := c.Param("id")
	if id != "" {
		if err1, userView := api.userService.Get(id); err1 != nil {
			global.Logger.Error("查询失败!", zap.Error(err1))
			response.OkWithMessage(err1.Error(), c)
			return
		} else {
			userInfoView.UserView = *userView
			err2, postIds := api.postService.SelectPostIdListByUserId(id)
			if err2 != nil {
				global.Logger.Error("查询失败!", zap.Error(err1))
				response.OkWithMessage(err1.Error(), c)
				return
			}
			userInfoView.PostIds = postIds
			var roleIds []string
			for _, roleView := range userView.Roles {
				if roleView.Id != constants.SYSTEM_ROLE_ADMIN_ID {
					roleIds = append(roleIds, roleView.Id)
				}
			}
			userInfoView.RoleIds = roleIds
		}
	}
	response.OkWithData(userInfoView, c)
}

// Page 分页获取User列表
// @Summary 分页获取User列表
// @Router /user/page [get]
func (api *UserApi) Page(c *gin.Context) {
	var pageInfo vo.UserPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, res := api.userService.Page(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// List 获取User列表
// @Summary 获取User列表
// @Router /user/list [get]
func (api *UserApi) List(c *gin.Context) {
	var view vo.UserView
	// 绑定查询参数到 view对象
	if err := c.ShouldBindQuery(&view); err != nil {
		response.FailWithMessage("获取参数解析失败!", c)
		return
	}
	// 判断是否需要根据用户获取数据
	// userId := api.GetLoginUserId(c)
	if err, res := api.userService.List(&view); err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// ResetPwd 重置密码
// @Summary 重置密码
// @Router /user/resetPwd [put]
func (api *UserApi) ResetPwd(c *gin.Context) {
	var req vo.UserView
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if constants.SYSTEM_ADMIN_ID == req.Id {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.userService.CheckUserDataScope(req.Id, api.GetLoginUserId(c)); err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	req.UpdateTime = utils.GetCurTimeStr()
	req.UpdateBy = api.GetLoginUserName(c)
	if err := api.userService.ResetPwd(&req); err != nil {
		response.FailWithMessage("重置密码失败", c)
		return
	}
	response.OkWithMessage("重置密码成功", c)
}

// ChangeStatus 更新状态
// @Summary 更新状态
// @Router /user/changeStatus [put]
func (api *UserApi) ChangeStatus(c *gin.Context) {
	var req vo.UserView
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if constants.SYSTEM_ADMIN_ID == req.Id {
		response.FailWithMessage("超级管理员不允许修改", c)
		return
	}
	if err := api.userService.CheckUserDataScope(req.Id, api.GetLoginUserId(c)); err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	req.UpdateTime = utils.GetCurTimeStr()
	req.UpdateBy = api.GetLoginUserName(c)
	if err := api.userService.ChangeStatus(&req); err != nil {
		response.FailWithMessage("更新状态失败", c)
		return
	}
	response.OkWithMessage("更新状态成功", c)
}

// GetAuthRole 根据用户id获取授权角色
// @Summary 根据用户id获取授权角色
// @Router /user/getAuthRole/{userId} [get]
func (api *UserApi) GetAuthRole(c *gin.Context) {
	userId := c.Param("id")
	err, userView := api.userService.Get(userId)
	if err != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	if err1, res := api.roleService.AssembleRolesByUserId(userId); err1 != nil {
		global.Logger.Error("获取数据失败!", zap.Error(err1))
		response.FailWithMessage("获取失败", c)
	} else {
		removeAdminRole(&res)
		response.OkWithData(gin.H{
			"user":  userView,
			"roles": res,
		}, c)
	}
}

// AuthRole 批量给用户授权角色
// @Summary 批量给用户授权角色
// @Router /user/authRole [put]
func (api *UserApi) AuthRole(c *gin.Context) {
	userId := c.Query("userId")
	roleIdStr := c.Query("roleIds")
	if userId == "" || roleIdStr == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	roleIds := strings.Split(roleIdStr, ",")
	if err := api.userService.CheckUserDataScope(userId, api.GetLoginUserId(c)); err != nil {
		response.FailWithMessage("没有权限访问用户数据", c)
		return
	}
	v := &vo.UserView{
		Id:         userId,
		RoleIds:    roleIds,
		UpdateTime: utils.GetCurTimeStr(),
		UpdateBy:   api.GetLoginUserName(c),
	}

	if err := api.userService.AuthRole(v); err != nil {
		response.FailWithMessage("授权失败", c)
		return
	}
	response.OkWithMessage("授权成功", c)
}

// 剔除超级管理员
func removeAdminRole(roles *[]*vo.RoleView) {
	for i := 0; i < len(*roles); i++ {
		if (*roles)[i].Id == constants.SYSTEM_ROLE_ADMIN_ID {
			*roles = append((*roles)[:i], (*roles)[i+1:]...)
			break
		}
	}
}

// SelectAllocatedList 查询已分配用户角色列表
// @Summary 查询已分配用户角色列表
// @Router /user/selectAllocatedList [get]
func (api *UserApi) SelectAllocatedList(c *gin.Context) {
	var pageInfo vo.UserPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, res := api.userService.SelectAllocatedList(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}

// SelectUnallocatedList 查询未分配用户角色列表
// @Summary 查询未分配用户角色列表
// @Router /user/selectUnallocatedList [get]
func (api *UserApi) SelectUnallocatedList(c *gin.Context) {
	var pageInfo vo.UserPageView
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage("获取分页数据解析失败!", c)
		return
	}
	user := api.GetLoginUser(c)
	if err, res := api.userService.SelectUnallocatedList(&pageInfo, user); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(res, "获取成功", c)
	}
}
