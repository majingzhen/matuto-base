// Package convert 自动生成模板 TableColumn
// @description <TODO description class purpose>
// @author
// @File: table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package vo

// TableColumnView 结构体
type TableColumnView struct {
	ColumnComment string `json:"columnComment" form:"columnComment"`
	ColumnName    string `json:"columnName" form:"columnName"`
	ColumnType    string `json:"columnType" form:"columnType"`
	DataType      string `json:"dataType" form:"dataType"`
	ColumnLength  int    `json:"columnLength" form:"columnLength"`
	CreateBy      string `json:"createBy" form:"createBy"`
	CreateTime    string `json:"createTime" form:"createTime"`
	DictType      string `json:"dictType" form:"dictType"`
	GoField       string `json:"goField" form:"goField"`
	GoType        string `json:"goType" form:"goType"`
	DefaultValue  string `json:"defaultValue" form:"defaultValue"`
	JsonField     string `json:"jsonField" form:"jsonField"`
	HtmlType      string `json:"htmlType" form:"htmlType"`
	ShowLabel     string `json:"showLabel" form:"showLabel"`
	Id            string `json:"id" form:"id"`
	IsEdit        string `json:"isEdit" form:"isEdit"`
	IsIncrement   string `json:"isIncrement" form:"isIncrement"`
	IsInsert      string `json:"isInsert" form:"isInsert"`
	IsList        string `json:"isList" form:"isList"`
	IsPk          string `json:"isPk" form:"isPk"`
	IsQuery       string `json:"isQuery" form:"isQuery"`
	IsRequired    string `json:"isRequired" form:"isRequired"`
	IsBase        string `json:"isBase" form:"isBase"`
	QueryType     string `json:"queryType" form:"queryType"`
	Sort          int    `json:"sort" form:"sort"`
	TableId       string `json:"tableId" form:"tableId"`
	UpdateBy      string `json:"updateBy" form:"updateBy"`
	UpdateTime    string `json:"updateTime" form:"updateTime"`
	Operation     string `json:"operation" form:"operation"`
}
