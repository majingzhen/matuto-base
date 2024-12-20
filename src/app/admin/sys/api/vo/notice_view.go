// Package convert 自动生成模板 Notice
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-14 14:18:43
package vo

import (
	"matuto-base/src/common"
)

// NoticeView 结构体
type NoticeView struct {
	Id            string `json:"id" form:"id"`
	NoticeTitle   string `json:"noticeTitle" form:"noticeTitle"`
	NoticeType    string `json:"noticeType" form:"noticeType"`
	NoticeContent string `json:"noticeContent" form:"noticeContent"`
	Status        string `json:"status" form:"status"`
	CreateBy      string `json:"createBy" form:"createBy"`
	CreateTime    string `json:"createTime" form:"createTime"`
	UpdateBy      string `json:"updateBy" form:"updateBy"`
	UpdateTime    string `json:"updateTime" form:"updateTime"`
	Remark        string `json:"remark" form:"remark"`
}

// NoticePageView 结构体
type NoticePageView struct {
	common.PageView

	NoticeTitle     string `json:"noticeTitle" form:"noticeTitle"`
	NoticeType      string `json:"noticeType" form:"noticeType"`
	NoticeContent   string `json:"noticeContent" form:"noticeContent"`
	Status          string `json:"status" form:"status"`
	CreateTime      string `json:"createTime" form:"createTime"`
	BeginUpdateTime string `json:"beginUpdateTime" form:"beginUpdateTime"`
	EndUpdateTime   string `json:"endUpdateTime" form:"endUpdateTime"`
}

// NoticeQueryView 结构体
type NoticeQueryView struct {
	NoticeTitle string `json:"noticeTitle" form:"noticeTitle"`

	NoticeType string `json:"noticeType" form:"noticeType"`

	NoticeContent string `json:"noticeContent" form:"noticeContent"`

	Status string `json:"status" form:"status"`

	CreateTime string `json:"createTime" form:"createTime"`

	BeginUpdateTime string `json:"beginUpdateTime" form:"beginUpdateTime"`
	EndUpdateTime   string `json:"endUpdateTime" form:"endUpdateTime"`
}

// NoticeCreateView 结构体
type NoticeCreateView struct {
	Id            string `json:"id" form:"id"`
	NoticeTitle   string `json:"noticeTitle" form:"noticeTitle"`
	NoticeType    string `json:"noticeType" form:"noticeType"`
	NoticeContent string `json:"noticeContent" form:"noticeContent"`
	Status        string `json:"status" form:"status"`
	CreateBy      string `json:"createBy" form:"createBy"`
	CreateTime    string `json:"createTime" form:"createTime"`
	UpdateBy      string `json:"updateBy" form:"updateBy"`
	UpdateTime    string `json:"updateTime" form:"updateTime"`
	Remark        string `json:"remark" form:"remark"`
}

// NoticeEditView 结构体
type NoticeEditView struct {
	Id string `json:"id" form:"id"`

	NoticeTitle string `json:"noticeTitle" form:"noticeTitle"`

	NoticeType string `json:"noticeType" form:"noticeType"`

	NoticeContent string `json:"noticeContent" form:"noticeContent"`

	Status string `json:"status" form:"status"`

	UpdateBy string `json:"updateBy" form:"updateBy"`

	UpdateTime string `json:"updateTime" form:"updateTime"`

	Remark string `json:"remark" form:"remark"`
}
