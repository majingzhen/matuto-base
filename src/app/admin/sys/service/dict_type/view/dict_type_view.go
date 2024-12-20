// Package convert 自动生成模板 DictType
// @description <TODO description class purpose>
// @author
// @File: dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import "matuto-base/src/common"

// DictTypeView 结构体

type DictTypeView struct {
	Id         string `json:"id"`
	DictName   string `json:"dictName"`
	DictType   string `json:"dictType"`
	Status     string `json:"status"`
	CreateBy   string `json:"createBy"`
	CreateTime string `json:"createTime"`
	UpdateBy   string `json:"updateBy"`
	UpdateTime string `json:"updateTime"`
	Remark     string `json:"remark"`
}

type DictTypePageView struct {
	common.PageView
	DictName   string `json:"dictName" form:"dictName"`
	DictType   string `json:"dictType" form:"dictType"`
	Status     string `json:"status" form:"status"`
	CreateTime string `json:"createTime" form:"createTime"`
}
