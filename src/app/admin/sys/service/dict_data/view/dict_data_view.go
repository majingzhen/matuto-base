// Package convert 自动生成模板 DictData
// @description <TODO description class purpose>
// @author
// @File: dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package view

import "matuto-base/src/common"

// DictDataView 结构体

type DictDataView struct {
	Id         string `json:"id" form:"id"`
	DictSort   int    `json:"dictSort" form:"dictSort"`
	DictLabel  string `json:"dictLabel" form:"dictLabel"`
	DictValue  string `json:"dictValue" form:"dictValue"`
	DictType   string `json:"dictType" form:"dictType"`
	CssClass   string `json:"cssClass" form:"cssClass"`
	ListClass  string `json:"listClass" form:"listClass"`
	IsDefault  string `json:"isDefault" form:"isDefault"`
	Status     string `json:"status" form:"status"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`
}

type DictDataPageView struct {
	common.PageView
	// TODO 按需修改
	Id         string `json:"id" form:"id"`
	DictSort   int    `json:"dictSort" form:"dictSort"`
	DictLabel  string `json:"dictLabel" form:"dictLabel"`
	DictValue  string `json:"dictValue" form:"dictValue"`
	DictType   string `json:"dictType" form:"dictType"`
	CssClass   string `json:"cssClass" form:"cssClass"`
	ListClass  string `json:"listClass" form:"listClass"`
	IsDefault  string `json:"isDefault" form:"isDefault"`
	Status     string `json:"status" form:"status"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`
}
