// Package api  UserApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"image/color"
	"matuto-base/src/app/admin/sys/service/menu"
	"matuto-base/src/app/admin/sys/service/role"
	"matuto-base/src/app/admin/sys/service/system/view"
	"matuto-base/src/app/admin/sys/service/user"
	"matuto-base/src/common/basic"
	"matuto-base/src/common/response"
	"matuto-base/src/framework"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type SystemApi struct {
	basic.BasicApi
	userService user.Service
	roleService role.RoleService
	menuService menu.MenuService
}

// store 验证码
var store = base64Captcha.DefaultMemStore

// Login 登录
// @Summary 登录系统
// @Router /sysOauth2/login [post]
func (api *SystemApi) Login(c *gin.Context) {
	var loginUserView view.LoginUserView
	_ = c.ShouldBindJSON(&loginUserView)
	// 校验验证码
	captcha := VerifyCaptcha(loginUserView.VerifyUuid, loginUserView.VerifyCode)
	if !captcha {
		response.FailWithMessage("验证码错误", c)
		return
	}
	if loginUserView.UserName == "" || loginUserView.Password == "" {
		global.Logger.Error("账号密码不能为空")
		response.FailWithMessage("账号密码不能为空", c)
		return
	}
	err, byUserName := api.userService.GetByUserName(loginUserView.UserName)
	if err != nil || byUserName == nil {
		global.Logger.Error("用户不存在", zap.Error(err))
		response.FailWithMessage("用户不存在", c)
		return
	}
	// 取加密密码
	hashedPassword := utils.EncryptionPassword(loginUserView.Password, byUserName.Salt)
	if hashedPassword != byUserName.Password {
		global.Logger.Error("登录失败")
		response.FailWithMessage("登录失败", c)
		return
	} else {
		// 判断是否分配角色
		_, userView := api.userService.Get(byUserName.Id)
		if userView.Roles == nil || len(userView.Roles) == 0 {
			global.Logger.Error("用户不存在", zap.Error(err))
			response.FailWithMessage("用户不存在", c)
			return
		}
		token, err := framework.GenerateToken(userView.Id, userView.UserName)
		if err != nil {
			response.FailWithMessage("登录失败", c)
			return
		}
		response.OkWithData(token, c)
	}
}

// GetUserInfo 获取用户信息
func (api *SystemApi) GetUserInfo(c *gin.Context) {
	userView := api.GetLoginUser(c)
	// 获取用户角色
	_, roles := api.roleService.GetRoleByUserId(userView)
	// 获取用户权限
	_, perms := api.menuService.GetMenuPermission(userView)
	// 获取用户菜单
	resView := view.LoginUserResView{
		UserInfo:    userView,
		Roles:       roles,
		Permissions: perms,
	}
	response.OkWithData(resView, c)
}

// GetRouters 获取路由信息
func (api *SystemApi) GetRouters(c *gin.Context) {
	userId := api.GetLoginUserId(c)
	err, tree := api.menuService.GetMenuTreeByUserId(userId)
	if err != nil {
		// 处理生成验证码时的错误
		response.FailWithMessage("获取路由失败", c)
	}
	response.OkWithData(tree, c)

}

// CaptchaImage 验证码
func (api *SystemApi) CaptchaImage(c *gin.Context) {
	//字符,公式,验证码配置
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          40,    //高度
		Width:           100,   //宽度
		NoiseCount:      0,     //干扰数
		ShowLineOptions: 2 | 4, //展示个数
		Length:          4,     //长度
		// Source:          "1234567890qwertyuiplkjhgfdsazxcvbnm", //验证码随机字符串来源
		Source: "1234567890", //验证码随机字符串来源
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, // 字体
	}
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		// 处理生成验证码时的错误
		response.FailWithMessage("登录失败", c)
	}
	response.OkWithData(&view.Captcha{
		Key: id,
		Img: b64s,
	}, c)
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	// 参数说明: id 验证码id, verifyValue 验证码的值, true: 验证成功后是否删除原来的验证码
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}

// Logout 退出登录
func (api *SystemApi) Logout(c *gin.Context) {
	response.OkWithMessage("success", c)
}
