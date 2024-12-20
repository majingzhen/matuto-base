// Package convert 自动生成模板 Dept
// @description <TODO description class purpose>
// @author
// @File: dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package vo

// DeptView 结构体

type DeptView struct {
	Id         string      `json:"id" form:"id"`
	ParentId   string      `json:"parentId" form:"parentId"`
	Ancestors  string      `json:"ancestors" form:"ancestors"`
	DeptName   string      `json:"deptName" form:"deptName"`
	OrderNum   int         `json:"orderNum" form:"orderNum"`
	Leader     string      `json:"leader" form:"leader"`
	Phone      string      `json:"phone" form:"phone"`
	Email      string      `json:"email" form:"email"`
	Status     string      `json:"status" form:"status"`
	CreateBy   string      `json:"createBy" form:"createBy"`
	CreateTime string      `json:"createTime" form:"createTime"`
	UpdateBy   string      `json:"updateBy" form:"updateBy"`
	UpdateTime string      `json:"updateTime" form:"updateTime"`
	Children   *[]DeptView `json:"children" form:"children"`
}

type DeptTreeView struct {
	Id       string          `json:"id" form:"id"`
	DeptName string          `json:"deptName" form:"deptName"`
	ParentId string          `json:"parentId" form:"parentId"`
	Children []*DeptTreeView `json:"children" form:"children"`
}
