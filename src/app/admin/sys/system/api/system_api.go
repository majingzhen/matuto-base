// Package api  SysUserApi 自动生成模板
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 13:41:26
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"image/color"
	"manager-gin/src/app/admin/sys/sys_user/service"
	"manager-gin/src/app/admin/sys/system/api/view"
	"manager-gin/src/common"
	"manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SystemApi struct {
}

var sysUserService = service.SysUserServiceApp

// store 验证码
var store = base64Captcha.DefaultMemStore

// Login 登录
// @Summary 登录系统
// @Router /sysOauth2/login [post]
func (systemApi *SystemApi) Login(c *gin.Context) {
	var loginUserView view.LoginUserView
	_ = c.ShouldBindJSON(&loginUserView)
	// 校验验证码
	captcha := VerifyCaptcha(loginUserView.VerifyKey, loginUserView.VerifyCode)
	if !captcha {
		response.FailWithMessage("验证码错误", c)
		return
	}
	if loginUserView.UserName == "" || loginUserView.Password == "" {
		global.Logger.Error("账号密码不能为空")
		response.FailWithMessage("账号密码不能为空", c)
		return
	}
	err, userView := sysUserService.GetByUserName(loginUserView.UserName)
	if err != nil || userView == nil {
		global.Logger.Error("用户名不存在", zap.Error(err))
		response.FailWithMessage("用户名不存在", c)
		return
	}
	// 取加密密码
	hashedPassword := utils.EncryptionPassword(loginUserView.Password, userView.Salt)
	if hashedPassword != userView.Password {
		global.Logger.Error("登录失败")
		response.FailWithMessage("登录失败", c)
		return
	} else {
		token, err := common.GenerateToken(userView.Id)
		if err != nil {
			response.FailWithMessage("登录失败", c)
			return
		}
		// 获取用户信息
		_, sysUserView := sysUserService.Get(userView.Id)
		resUserView := view.SysUserInfoView{
			Id:          sysUserView.Id,
			UserName:    sysUserView.UserName,
			UserType:    sysUserView.UserType,
			OrgId:       sysUserView.OrgId,
			NickName:    sysUserView.NickName,
			Email:       sysUserView.Email,
			PhoneNumber: sysUserView.PhoneNumber,
			Sex:         sysUserView.Sex,
			Avatar:      sysUserView.Avatar,
		}
		// 获取用户菜单
		resView := view.LoginUserResView{Token: token, SysUserInfo: resUserView}
		response.OkWithData(resView, c)
	}
}

// CaptchaImage 验证码
func (systemApi *SystemApi) CaptchaImage(c *gin.Context) {
	//字符,公式,验证码配置
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          40,                                     //高度
		Width:           100,                                    //宽度
		NoiseCount:      0,                                      //干扰数
		ShowLineOptions: 2 | 4,                                  //展示个数
		Length:          4,                                      //长度
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm", //验证码随机字符串来源
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
	response.OkWithData(common.Captcha{
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
