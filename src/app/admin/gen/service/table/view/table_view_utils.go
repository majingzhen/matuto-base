// Package convert
// @description <TODO description class purpose>
// @author
// @File: gen_table
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

type TableViewUtils struct{}

func (viewUtils *TableViewUtils) Data2View(data *model.Table) (err error, view *TableView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableViewUtils View2Data error: %v", e)
			global.Logger.Error("TableViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp TableView
	tmp.BusinessName = data.BusinessName
	tmp.CreateBy = data.CreateBy
	tmp.CreateTime = utils.Time2Str(data.CreateTime)
	tmp.FunctionAuthor = data.FunctionAuthor
	tmp.FunctionName = data.FunctionName
	tmp.GenPath = data.GenPath
	tmp.GenType = data.GenType
	tmp.Id = data.Id
	tmp.ModuleName = data.ModuleName
	tmp.Options = data.Options
	tmp.PackageName = data.PackageName
	tmp.Remark = data.Remark
	tmp.StructName = data.StructName
	tmp.SubTableFkName = data.SubTableFkName
	tmp.SubTableName = data.SubTableName
	tmp.TableComment = data.TableComment
	tmp.TableName = data.Name
	tmp.TplCategory = data.TplCategory
	tmp.UpdateBy = data.UpdateBy
	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)
	view = &tmp
	return
}
func (viewUtils *TableViewUtils) View2Data(view *TableView) (err error, data *model.Table) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableViewUtils View2Data error: %v", e)
			global.Logger.Error("TableViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.Table
	tmp.BusinessName = view.BusinessName
	tmp.CreateBy = view.CreateBy
	tmp.CreateTime = utils.Str2Time(view.CreateTime)
	tmp.FunctionAuthor = view.FunctionAuthor
	tmp.FunctionName = view.FunctionName
	tmp.GenType = view.GenType
	tmp.GenPath = view.GenPath
	tmp.Id = view.Id
	tmp.ModuleName = view.ModuleName
	tmp.Options = view.Options
	tmp.PackageName = view.PackageName
	tmp.Remark = view.Remark
	tmp.StructName = view.StructName
	tmp.SubTableFkName = view.SubTableFkName
	tmp.SubTableName = view.SubTableName
	tmp.TableComment = view.TableComment
	tmp.Name = view.TableName
	tmp.TplCategory = view.TplCategory
	tmp.UpdateBy = view.UpdateBy
	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)
	data = &tmp
	return
}

func (viewUtils *TableViewUtils) View2DataList(viewList []*TableView) (err error, dataList []*model.Table) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableViewUtils View2DataList error: %v", e)
			global.Logger.Error("TableViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.Table
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

func (viewUtils *TableViewUtils) Data2ViewList(dataList []*model.Table) (err error, viewList []*TableView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("TableViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*TableView
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

func (viewUtils *TableViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("TableViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("TableViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.Table); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
