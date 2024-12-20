// Package convert 自动生成模板 Menu
// @description <TODO description class purpose>
// @author
// @File: menu
// @version 1.0.0
// @create 2023-08-20 21:21:34
package vo

// MenuView 结构体
type MenuView struct {
	Id         string `json:"id" form:"id"`
	MenuName   string `json:"menuName" form:"menuName"`
	ParentId   string `json:"parentId" form:"parentId"`
	OrderNum   int    `json:"orderNum" form:"orderNum"`
	Path       string `json:"path" form:"path"`
	Component  string `json:"component" form:"component"`
	Query      string `json:"query" form:"query"`
	IsFrame    int    `json:"isFrame" form:"isFrame"`
	IsCache    int    `json:"isCache" form:"isCache"`
	MenuType   string `json:"menuType" form:"menuType"`
	Visible    string `json:"visible" form:"visible"`
	Status     string `json:"status" form:"status"`
	Perms      string `json:"perms" form:"perms"`
	Icon       string `json:"icon" form:"icon"`
	CreateBy   string `json:"createBy" form:"createBy"`
	CreateTime string `json:"createTime" form:"createTime"`
	UpdateBy   string `json:"updateBy" form:"updateBy"`
	UpdateTime string `json:"updateTime" form:"updateTime"`
	Remark     string `json:"remark" form:"remark"`
}

type MenuTree struct {
	Id       string      `json:"id" form:"id"`
	Label    string      `json:"label" form:"label"`
	Children []*MenuTree `json:"children" form:"children"`
}

type RouterView struct {
	Name       string        `json:"name"`
	Path       string        `json:"path"`
	Hidden     bool          `json:"hidden"`
	Redirect   string        `json:"redirect"`
	Component  string        `json:"component"`
	Query      string        `json:"query"`
	AlwaysShow bool          `json:"alwaysShow"`
	Meta       *MetaView     `json:"meta"`
	Children   []*RouterView `json:"children"`
}

type MetaView struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	NoCache bool   `json:"NoCache"`
	Link    string `json:"link"`
}
