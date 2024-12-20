// Package convert 自动生成模板 User
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-21 14:20:37
package vo

import (
	"matuto-base/src/common"
)

// UserView 结构体

type UserView struct {
	Id          string      `json:"id" form:"id"`
	DeptId      string      `json:"deptId" form:"deptId"`
	UserName    string      `json:"userName" form:"userName"`
	NickName    string      `json:"nickName" form:"nickName"`
	UserType    string      `json:"userType" form:"userType"`
	Email       string      `json:"email" form:"email"`
	PhoneNumber string      `json:"phoneNumber" form:"phoneNumber"`
	Sex         string      `json:"sex" form:"sex"`
	Avatar      string      `json:"avatar" form:"avatar"`
	Password    string      `json:"password" form:"password"`
	Salt        string      `json:"salt" form:"salt"`
	Status      string      `json:"status" form:"status"`
	LoginIp     string      `json:"loginIp" form:"loginIp"`
	LoginDate   string      `json:"loginDate" form:"loginDate"`
	CreateBy    string      `json:"createBy" form:"createBy"`
	CreateTime  string      `json:"createTime" form:"createTime"`
	UpdateBy    string      `json:"updateBy" form:"updateBy"`
	UpdateTime  string      `json:"updateTime" form:"updateTime"`
	Remark      string      `json:"remark" form:"remark"`
	Roles       []*RoleView `json:"roles"`
	Dept        *DeptView   `json:"dept"`
	RoleIds     []string    `json:"roleIds" form:"roleIds"`
	PostIds     []string    `json:"postIds" form:"postIds"`
}

type UserPageView struct {
	common.PageView
	// TODO 按需修改
	Id           string `json:"id" form:"id"`
	DeptId       string `json:"deptId" form:"deptId"`
	UserName     string `json:"userName" form:"userName"`
	NickName     string `json:"nickName" form:"nickName"`
	UserType     string `json:"userType" form:"userType"`
	Email        string `json:"email" form:"email"`
	PhoneNumber  string `json:"phoneNumber" form:"phoneNumber"`
	Sex          string `json:"sex" form:"sex"`
	Status       string `json:"status" form:"status"`
	CreateTime   string `json:"createTime" form:"createTime"`
	RoleId       string `json:"roleId" form:"roleId"`
	DataScopeSql string
}

// UserInfoView 结构体
type UserInfoView struct {
	UserView
	RoleIds []string    `json:"roleIds" form:"roleIds"`
	Roles   []*RoleView `json:"roles" form:"roles"`
	PostIds []string    `json:"postIds" form:"postIds"`
	Posts   []*PostView `json:"posts" form:"posts"`
}
