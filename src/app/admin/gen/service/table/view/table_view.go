// Package convert 自动生成模板 Table
// @description <TODO description class purpose>
// @author
// @File: table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package view

import (
	"matuto-base/src/app/admin/gen/service/table_column/view"
	"matuto-base/src/common"
)

// TableView 结构体
type TableView struct {
	BusinessName   string                  `json:"businessName" form:"businessName"`
	CreateBy       string                  `json:"createBy" form:"createBy"`
	CreateTime     string                  `json:"createTime" form:"createTime"`
	FunctionAuthor string                  `json:"functionAuthor" form:"functionAuthor"`
	FunctionName   string                  `json:"functionName" form:"functionName"`
	GenPath        string                  `json:"genPath" form:"genPath"`
	GenType        string                  `json:"genType" form:"genType"`
	Id             string                  `json:"id" form:"id"`
	ModuleName     string                  `json:"moduleName" form:"moduleName"`
	Options        string                  `json:"options" form:"options"`
	PackageName    string                  `json:"packageName" form:"packageName"`
	Remark         string                  `json:"remark" form:"remark"`
	StructName     string                  `json:"structName" form:"structName"`
	SubTableFkName string                  `json:"subTableFkName" form:"subTableFkName"`
	SubTableName   string                  `json:"subTableName" form:"subTableName"`
	TableComment   string                  `json:"tableComment" form:"tableComment"`
	TableName      string                  `json:"tableName" form:"tableName"`
	TplCategory    string                  `json:"tplCategory" form:"tplCategory"`
	UpdateBy       string                  `json:"updateBy" form:"updateBy"`
	UpdateTime     string                  `json:"updateTime" form:"updateTime"`
	ColumnList     []*view.TableColumnView `json:"columnList" form:"columnList"`
	TableViewOptions
	Author       string
	PKColumn     *view.TableColumnView
	SearchColumn []*view.TableColumnView
	InsertColumn []*view.TableColumnView
	EditColumn   []*view.TableColumnView
	ListColumn   []*view.TableColumnView
	Dicts        string
}

// TableViewOptions Options 扩展字段
type TableViewOptions struct {
	TreeCode       string `json:"treeCode" form:"treeCode" gorm:"column:tree_code;comment:树编码;"`
	TreeParentCode string `json:"treeParentCode" form:"treeParentCode" gorm:"column:tree_parent_code;comment:树父编码;"`
	TreeName       string `json:"treeName" form:"treeName" gorm:"column:tree_name;comment:树名称;"`
	ParentMenuName string `json:"parentMenuName" form:"parentMenuName" gorm:"column:parent_menu_name;comment:父菜单名称;"`
	ParentMenuId   string `json:"parentMenuId" form:"parentMenuId" gorm:"column:parent_menu_id;comment:父菜单编号;"`
}

type TablePageView struct {
	common.PageView
	// TODO 按需修改
	BusinessName   string `json:"businessName" form:"businessName"`
	CreateBy       string `json:"createBy" form:"createBy"`
	CreateTime     string `json:"createTime" form:"createTime"`
	FunctionAuthor string `json:"functionAuthor" form:"functionAuthor"`
	FunctionName   string `json:"functionName" form:"functionName"`
	Path           string `json:"genPath" form:"genPath"`
	Type           string `json:"genType" form:"genType"`
	Id             string `json:"id" form:"id"`
	ModuleName     string `json:"moduleName" form:"moduleName"`
	Options        string `json:"options" form:"options"`
	PackageName    string `json:"packageName" form:"packageName"`
	Remark         string `json:"remark" form:"remark"`
	StructName     string `json:"structName" form:"structName"`
	SubTableFkName string `json:"subTableFkName" form:"subTableFkName"`
	SubTableName   string `json:"subTableName" form:"subTableName"`
	TableComment   string `json:"tableComment" form:"tableComment"`
	TableName      string `json:"tableName" form:"tableName"`
	TplCategory    string `json:"tplCategory" form:"tplCategory"`
	UpdateBy       string `json:"updateBy" form:"updateBy"`
	UpdateTime     string `json:"updateTime" form:"updateTime"`
}

type TableQueryView struct {
	// TODO 按需修改
	BusinessName   string `json:"businessName" form:"businessName"`
	CreateBy       string `json:"createBy" form:"createBy"`
	CreateTime     string `json:"createTime" form:"createTime"`
	FunctionAuthor string `json:"functionAuthor" form:"functionAuthor"`
	FunctionName   string `json:"functionName" form:"functionName"`
	Path           string `json:"genPath" form:"genPath"`
	Type           string `json:"genType" form:"genType"`
	Id             string `json:"id" form:"id"`
	ModuleName     string `json:"moduleName" form:"moduleName"`
	Options        string `json:"options" form:"options"`
	PackageName    string `json:"packageName" form:"packageName"`
	Remark         string `json:"remark" form:"remark"`
	StructName     string `json:"structName" form:"structName"`
	SubTableFkName string `json:"subTableFkName" form:"subTableFkName"`
	SubTableName   string `json:"subTableName" form:"subTableName"`
	TableComment   string `json:"tableComment" form:"tableComment"`
	TableName      string `json:"tableName" form:"tableName"`
	TplCategory    string `json:"tplCategory" form:"tplCategory"`
	UpdateBy       string `json:"updateBy" form:"updateBy"`
	UpdateTime     string `json:"updateTime" form:"updateTime"`
}

// GenTableAndColumnView 结构体
type GenTableAndColumnView struct {
	BusinessName   string                  `json:"businessName" form:"businessName"`
	CreateBy       string                  `json:"createBy" form:"createBy"`
	CreateTime     string                  `json:"createTime" form:"createTime"`
	FunctionAuthor string                  `json:"functionAuthor" form:"functionAuthor"`
	FunctionName   string                  `json:"functionName" form:"functionName"`
	GenPath        string                  `json:"genPath" form:"genPath"`
	GenType        string                  `json:"genType" form:"genType"`
	Id             string                  `json:"id" form:"id"`
	ModuleName     string                  `json:"moduleName" form:"moduleName"`
	Options        string                  `json:"options" form:"options"`
	PackageName    string                  `json:"packageName" form:"packageName"`
	Remark         string                  `json:"remark" form:"remark"`
	StructName     string                  `json:"structName" form:"structName"`
	SubTableFkName string                  `json:"subTableFkName" form:"subTableFkName"`
	SubTableName   string                  `json:"subTableName" form:"subTableName"`
	TableComment   string                  `json:"tableComment" form:"tableComment"`
	TableName      string                  `json:"tableName" form:"tableName"`
	TplCategory    string                  `json:"tplCategory" form:"tplCategory"`
	UpdateBy       string                  `json:"updateBy" form:"updateBy"`
	UpdateTime     string                  `json:"updateTime" form:"updateTime"`
	ColumnList     []*view.TableColumnView `json:"columnList" form:"columnList"`
	Author         string                  `json:"author" form:"author"`
	TableViewOptions
}
