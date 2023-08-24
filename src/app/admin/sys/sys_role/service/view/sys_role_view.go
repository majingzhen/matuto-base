// Package view 自动生成模板 SysRole
// @description <TODO description class purpose>
// @author
// @File: sys_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package view

// SysRoleView 结构体
type SysRoleView struct {
	Id                string   `json:"id" form:"id"`
	RoleName          string   `json:"roleName" form:"roleName"`
	RoleKey           string   `json:"roleKey" form:"roleKey"`
	RoleSort          int      `json:"roleSort" form:"roleSort"`
	DataScope         string   `json:"dataScope" form:"dataScope"`
	MenuCheckStrictly bool     `json:"menuCheckStrictly" form:"menuCheckStrictly"`
	DeptCheckStrictly bool     `json:"deptCheckStrictly" form:"deptCheckStrictly"`
	Status            string   `json:"status" form:"status"`
	CreateBy          string   `json:"createBy" form:"createBy"`
	CreateTime        string   `json:"createTime" form:"createTime"`
	UpdateBy          string   `json:"updateBy" form:"updateBy"`
	UpdateTime        string   `json:"updateTime" form:"updateTime"`
	Remark            string   `json:"remark" form:"remark"`
	Permissions       []string `json:"permissions"`
	MenuIds           []string `json:"menuIds"`
	DeptIds           []string `json:"deptIds"`
	Flag              bool     `json:"flag"`
}

type SysRolePageView struct {
	// TODO 按需修改
	Id         string `json:"id" form:"id"`
	RoleName   string `json:"roleName" form:"roleName"`
	RoleKey    string `json:"roleKey" form:"roleKey"`
	RoleSort   int    `json:"roleSort" form:"roleSort"`
	DataScope  string `json:"dataScope" form:"dataScope"`
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

type SysUserRoleView struct {
	UserId string `json:"userId" form:"userId"`
	RoleId string `json:"roleId" form:"roleId"`
}
