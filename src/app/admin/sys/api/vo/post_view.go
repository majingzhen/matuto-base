// Package convert 自动生成模板 Post
// @description <TODO description class purpose>
// @author
// @File: post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package vo

import "matuto-base/src/common"

// PostView 结构体

type PostView struct {
	Id         string `json:"id" form:"id"`
	PostCode   string `json:"postCode" form:"postCode"`
	PostName   string `json:"postName" form:"postName"`
	PostSort   int    `json:"postSort" form:"postSort"`
	Status     string `json:"status" form:"status"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`
}

type PostPageView struct {
	common.PageView
	// TODO 按需修改
	Id         string `json:"id" form:"id"`
	PostCode   string `json:"postCode" form:"postCode"`
	PostName   string `json:"postName" form:"postName"`
	PostSort   int    `json:"postSort" form:"postSort"`
	Status     string `json:"status" form:"status"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`
}
