// Package view 自动生成模板 SysPost
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package view

// SysPostView 结构体

type SysPostView struct {
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

type SysPostPageView struct {
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

	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string `json:"isAsc" form:"isAsc"`                 //排序方式
	PageNum       int    `json:"pageNum" form:"pageNum"`             //当前页码
	PageSize      int    `json:"pageSize" form:"pageSize"`           //每页数
}