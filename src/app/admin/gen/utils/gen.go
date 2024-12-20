package utils

import (
	"matuto-base/src/app/admin/gen/model"
	"matuto-base/src/app/admin/gen/service/table/view"
	"matuto-base/src/common/constants"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strconv"
	"strings"
)

// InitTable 初始化表结构体
func InitTable(genTable *model.Table, operName string) *model.Table {
	genTable.CreateBy = operName
	genTable.StructName = convertStructName(genTable.Name)
	genTable.PackageName = global.Viper.GetString("gen.package_name")
	genTable.ModuleName = genModuleName(global.Viper.GetString("gen.package_name"))
	genTable.BusinessName = getBusinessName(genTable.Name)
	genTable.FunctionName = genTable.TableComment
	genTable.FunctionAuthor = global.Viper.GetString("gen.author")
	genTable.TableComment = strings.Replace(genTable.TableComment, "表", "", 1)
	return genTable
}

// convertStructName 转换struct名称
func convertStructName(tableName string) string {
	autoRemovePre := global.Viper.GetBool("gen.auto_remove_pre")
	tablePrefix := global.Viper.GetString("gen.table_prefix")
	if autoRemovePre && tablePrefix != "" {
		searchList := strings.Split(tablePrefix, ",")
		tableName = replaceFirst(tableName, searchList)
	}
	return utils.ToTitle(tableName)
}

// replaceFirst 批量替换前缀
func replaceFirst(replaceMen string, searchList []string) string {
	text := replaceMen
	for _, s := range searchList {
		if strings.HasPrefix(text, s) {
			text = strings.Replace(replaceMen, s, "", 1)
			break
		}
	}
	return text
}

// InitColumnField 初始化列属性字段
func InitColumnField(column *model.TableColumn, table *model.Table) *model.TableColumn {
	dataType := getColumnType(column.ColumnType)
	columnName := column.ColumnName
	column.TableId = table.Id
	column.CreateBy = table.CreateBy
	column.CreateTime = utils.GetCurTime()
	// 设置结构体字段名
	column.GoField = utils.ToTitle(columnName)
	column.JsonField = utils.ToCamelCase(columnName)
	column.ShowLabel = clearBracket(column.ColumnComment)
	column.GoType = constants.TYPE_INTERFACE
	column.DefaultValue = constants.DEFAULT_INTERFACE
	column.QueryType = constants.QUERY_EQ
	if utils.ContainsStr(constants.COLUMN_TYPE_STR, dataType) || utils.ContainsStr(constants.COLUMN_TYPE_TEXT, dataType) {
		columnLength := getColumnLength(column.ColumnType)
		if columnLength >= 500 || utils.ContainsStr(constants.COLUMN_TYPE_TEXT, dataType) {
			column.HtmlType = constants.HTML_TEXTAREA
			column.GoType = constants.TYPE_BYTE_SLICE
			column.DefaultValue = constants.DEFAULT_STR
		} else {
			column.HtmlType = constants.HTML_INPUT
			column.GoType = constants.TYPE_STRING
			column.DefaultValue = constants.DEFAULT_STR
		}
	} else if utils.ContainsStr(constants.COLUMN_TYPE_TIME, dataType) {
		column.HtmlType = constants.HTML_DATETIME
		column.GoType = constants.TYPE_DATE
		column.DefaultValue = constants.DEFAULT_INTERFACE
	} else if utils.ContainsStr(constants.COLUMN_TYPE_NUMBER, dataType) {
		column.HtmlType = constants.HTML_INPUT
		column.GoType = constants.TYPE_INTEGER
		column.DefaultValue = constants.DEFAULT_NUM
	} else if utils.ContainsStr(constants.COLUMN_TYPE_FLOAT, dataType) {
		column.HtmlType = constants.HTML_INPUT
		column.GoType = constants.TYPE_FLOAT
		column.DefaultValue = constants.DEFAULT_NUM
	}
	// 插入字段
	column.IsInsert = constants.REQUIRE

	// 编辑字段
	if !utils.ContainsStr(constants.COLUMN_NAME_NOT_EDIT, columnName) && column.IsPk != "1" {
		column.IsEdit = constants.REQUIRE
	}
	// 列表字段
	if !utils.ContainsStr(constants.COLUMN_NAME_NOT_LIST, columnName) && column.IsPk != "1" {
		column.IsList = constants.REQUIRE
	}
	// 查询字段
	if !utils.ContainsStr(constants.COLUMN_NAME_NOT_QUERY, columnName) && column.IsPk != "1" {
		column.IsQuery = constants.REQUIRE
	}
	// 是否为基础列
	if utils.ContainsStr(constants.BASE_ENTITY, columnName) || column.IsPk == "1" {
		column.IsBase = constants.REQUIRE
	}
	// 状态字段设置单选框
	if utils.EndsWithIgnoreCase(columnName, "status") || utils.EndsWithIgnoreCase(columnName, "flag") || utils.BeginsWithIgnoreCase(columnName, "is") {
		column.HtmlType = constants.HTML_RADIO
	} else if utils.EndsWithIgnoreCase(columnName, "type") || utils.EndsWithIgnoreCase(columnName, "sex") {
		// 类型&性别字段设置下拉框
		column.HtmlType = constants.HTML_SELECT
	} else if utils.EndsWithIgnoreCase(columnName, "image") {
		// 图片字段设置图片上传控件
		column.HtmlType = constants.HTML_IMAGE_UPLOAD
	} else if utils.EndsWithIgnoreCase(columnName, "file") {
		// 文件字段设置文件上传控件
		column.HtmlType = constants.HTML_FILE_UPLOAD
	} else if utils.EndsWithIgnoreCase(columnName, "content") {
		// 内容字段设置富文本控件
		column.HtmlType = constants.HTML_EDITOR
	}

	// 查询方式
	if column.IsQuery == constants.REQUIRE {
		if column.HtmlType == constants.HTML_DATETIME {
			column.QueryType = constants.QUERY_BETWEEN
		} else if column.HtmlType == constants.HTML_SELECT {
			column.QueryType = constants.QUERY_EQ
		} else {
			column.QueryType = constants.QUERY_LIKE
		}
	}
	return column
}

// clearBracket 清除括号内容
func clearBracket(comment string) string {
	eIndex := strings.Index(comment, "(")
	cIndex := strings.Index(comment, "（")
	if eIndex > 0 {
		comment = comment[:eIndex]
	}
	if cIndex > 0 {
		comment = comment[:cIndex]
	}
	return comment
}

// getColumnByTableName 根据表名获取列信息
func getColumnLength(columnType string) int {
	if strings.Contains(columnType, "(") {
		length := columnType[strings.Index(columnType, "(")+1 : strings.Index(columnType, ")")]
		res, _ := strconv.Atoi(length)
		return res
	} else {
		return 0
	}
}

// getColumnType 获取列类型
func getColumnType(columnType string) string {
	if strings.Contains(columnType, "(") {
		return columnType[:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

// getBusinessName 获取业务名称
func getBusinessName(name string) string {
	autoRemovePre := global.Viper.GetBool("gen.auto_remove_pre")
	tablePrefix := global.Viper.GetString("gen.table_prefix")
	if autoRemovePre && tablePrefix != "" {
		return strings.ReplaceAll(name, tablePrefix, "")
	} else {
		return name
	}
}

// genModuleName 获取模块名称
func genModuleName(packageName string) string {
	lastIndex := strings.LastIndex(packageName, "/")
	return packageName[lastIndex+1:]
}

// GenTemplatePath 获取模板路径
func GenTemplatePath(tplCategory string) []string {
	if tplCategory == "" {
		tplCategory = constants.TPL_CRUD
	}
	return []string{
		"./resources/tmpl/" + tplCategory + "/go/model.go.txt",
		"./resources/tmpl/" + tplCategory + "/go/dao.go.txt",
		"./resources/tmpl/" + tplCategory + "/go/service.go.txt",
		"./resources/tmpl/" + tplCategory + "/go/convert.go.txt",
		"./resources/tmpl/" + tplCategory + "/go/view_utils.go.txt",
		"./resources/tmpl/" + tplCategory + "/go/api.go.txt",
		"./resources/tmpl/" + tplCategory + "/go/router.go.txt",
		"./resources/tmpl/" + tplCategory + "/js/api.js.txt",
		"./resources/tmpl/" + tplCategory + "/vue/index.vue.txt",
	}
}

// GetDictList 获取字典列表
func GetDictList(table *view.TableView) []string {
	dicts := make([]string, 0)
	if table != nil && table.ColumnList != nil {
		for _, columnView := range table.ColumnList {
			if (!utils.ContainsStr(constants.BASE_ENTITY, columnView.JsonField) && !utils.ContainsStr(constants.TREE_ENTITY, columnView.ColumnName)) && columnView.DictType != "" && utils.ContainsStr(constants.DICT_HTML_TYPE, columnView.HtmlType) {
				dicts = append(dicts, "'"+columnView.DictType+"'")
			}
		}
	}
	return dicts
}
