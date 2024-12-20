// Package convert 自动生成模板 User
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"matuto-base/src/app/admin/sys/service/user/view"
)

// LoginUserView 结构体
type LoginUserView struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	VerifyCode string `json:"code"`
	VerifyUuid string `json:"uuid"`
}

type LoginUserResView struct {
	UserInfo    *view.UserView `json:"user"`
	Roles       []string       `json:"roles"`
	Permissions []string       `json:"permissions"`
}

// Captcha 验证码响应
type Captcha struct {
	Img interface{} `json:"img"` //数据内容
	Key string      `json:"key"` //验证码ID
}
