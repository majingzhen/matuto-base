// Package convert
// @description <TODO description class purpose>
// @author
// @File: gen_table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/gen/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type TableColumnViewUtils struct{}

func (viewUtils *TableColumnViewUtils) Data2View(data *model.TableColumn) (err error, view *TableColumnView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableColumnViewUtils View2Data error: %v", e)
			global.Logger.Error("TableColumnViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp TableColumnView

	tmp.ColumnComment = data.ColumnComment

	tmp.ColumnName = data.ColumnName

	tmp.ColumnType = data.ColumnType

	tmp.ColumnLength = data.ColumnLength

	tmp.DataType = data.DataType

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.ShowLabel = data.ShowLabel

	tmp.DictType = data.DictType

	tmp.GoField = data.GoField

	tmp.GoType = data.GoType

	tmp.DefaultValue = data.DefaultValue

	tmp.JsonField = data.JsonField

	tmp.HtmlType = data.HtmlType

	tmp.Id = data.Id

	tmp.IsEdit = data.IsEdit

	tmp.IsIncrement = data.IsIncrement

	tmp.IsInsert = data.IsInsert

	tmp.IsList = data.IsList

	tmp.IsPk = data.IsPk

	tmp.IsQuery = data.IsQuery

	tmp.IsRequired = data.IsRequired

	tmp.IsBase = data.IsBase

	tmp.QueryType = data.QueryType

	tmp.Sort = data.Sort

	tmp.TableId = data.TableId

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)
	view = &tmp
	return
}
func (viewUtils *TableColumnViewUtils) View2Data(view *TableColumnView) (err error, data *model.TableColumn) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableColumnViewUtils View2Data error: %v", e)
			global.Logger.Error("TableColumnViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.TableColumn

	tmp.ColumnComment = view.ColumnComment

	tmp.ColumnName = view.ColumnName

	tmp.ColumnType = view.ColumnType

	tmp.ColumnLength = view.ColumnLength

	tmp.DataType = view.DataType

	tmp.ShowLabel = view.ShowLabel

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.DictType = view.DictType

	tmp.GoField = view.GoField

	tmp.GoType = view.GoType

	tmp.DefaultValue = view.DefaultValue

	tmp.JsonField = view.JsonField

	tmp.HtmlType = view.HtmlType

	tmp.Id = view.Id
	tmp.IsEdit = view.IsEdit
	tmp.IsIncrement = view.IsIncrement
	tmp.IsInsert = view.IsInsert
	tmp.IsList = view.IsList
	tmp.IsPk = view.IsPk
	tmp.IsQuery = view.IsQuery
	tmp.IsRequired = view.IsRequired
	tmp.IsBase = view.IsBase
	tmp.QueryType = view.QueryType
	tmp.Sort = view.Sort
	tmp.TableId = view.TableId
	tmp.UpdateBy = view.UpdateBy
	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)
	data = &tmp
	return
}

func (viewUtils *TableColumnViewUtils) View2DataList(viewList []*TableColumnView) (err error, dataList []*model.TableColumn) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableColumnViewUtils View2DataList error: %v", e)
			global.Logger.Error("TableColumnViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.TableColumn
		for i := range viewList {
			view := viewList[i]
			err, data := viewUtils.View2Data(view)
			if err == nil {
				dataTmpList = append(dataTmpList, data)
			}
		}
		dataList = dataTmpList
	}
	return
}

func (viewUtils *TableColumnViewUtils) Data2ViewList(dataList []*model.TableColumn) (err error, viewList []*TableColumnView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableColumnViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("TableColumnViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*TableColumnView
		for i := range dataList {
			data := dataList[i]
			err, view := viewUtils.Data2View(data)
			if err == nil {
				viewTmpList = append(viewTmpList, view)
			}
		}
		viewList = viewTmpList
	}
	return
}

func (viewUtils *TableColumnViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableColumnViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("TableColumnViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.TableColumn); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
