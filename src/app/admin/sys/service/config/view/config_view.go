// Package convert 自动生成模板 Config
// @description <TODO description class purpose>
// @author
// @File: config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package view

import "matuto-base/src/common"

// ConfigView 结构体

type ConfigView struct {
	Id          string `json:"id" form:"id"`
	ConfigName  string `json:"configName" form:"configName"`
	ConfigKey   string `json:"configKey" form:"configKey"`
	ConfigValue string `json:"configValue" form:"configValue"`
	ConfigType  string `json:"configType" form:"configType"`
	CreateBy    string `json:"createBy" form:"createBy"`
	CreateTime  string `json:"createTime" form:"createTime"`
	UpdateBy    string `json:"updateBy" form:"updateBy"`
	UpdateTime  string `json:"updateTime" form:"updateTime"`
	Remark      string `json:"remark" form:"remark"`
}

type ConfigPageView struct {
	common.PageView
	// TODO 按需修改
	Id          string `json:"id" form:"id"`
	ConfigName  string `json:"configName" form:"configName"`
	ConfigKey   string `json:"configKey" form:"configKey"`
	ConfigValue string `json:"configValue" form:"configValue"`
	ConfigType  string `json:"configType" form:"configType"`
	CreateBy    string `json:"createBy" form:"createBy"`
	CreateTime  string `json:"createTime" form:"createTime"`
	UpdateBy    string `json:"updateBy" form:"updateBy"`
	UpdateTime  string `json:"updateTime" form:"updateTime"`
	Remark      string `json:"remark" form:"remark"`
}
