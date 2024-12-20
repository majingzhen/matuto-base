// Package convert
// @description <TODO description class purpose>
// @author
// @File: dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type DictDataViewUtils struct{}

func (viewUtils *DictDataViewUtils) Data2View(data *model.DictData) (err error, view *DictDataView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictDataViewUtils View2Data error: %v", e)
			global.Logger.Error("DictDataViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp DictDataView

	tmp.Id = data.Id

	tmp.DictSort = data.DictSort

	tmp.DictLabel = data.DictLabel

	tmp.DictValue = data.DictValue

	tmp.DictType = data.DictType

	tmp.CssClass = data.CssClass

	tmp.ListClass = data.ListClass

	tmp.IsDefault = data.IsDefault

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *DictDataViewUtils) View2Data(view *DictDataView) (err error, data *model.DictData) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictDataViewUtils View2Data error: %v", e)
			global.Logger.Error("DictDataViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.DictData

	tmp.Id = view.Id

	tmp.DictSort = view.DictSort

	tmp.DictLabel = view.DictLabel

	tmp.DictValue = view.DictValue

	tmp.DictType = view.DictType

	tmp.CssClass = view.CssClass

	tmp.ListClass = view.ListClass

	tmp.IsDefault = view.IsDefault

	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *DictDataViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictDataViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("DictDataViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.DictData); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}

func (viewUtils *DictDataViewUtils) View2DataList(viewList []*DictDataView) (err error, dataList []*model.DictData) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictDataViewUtils View2DataList error: %v", e)
			global.Logger.Error("DictDataViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.DictData
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

func (viewUtils *DictDataViewUtils) Data2ViewList(dataList []*model.DictData) (err error, viewList []*DictDataView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictDataViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("DictDataViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*DictDataView
		for i := range dataList {
			data := (dataList)[i]
			err, view := viewUtils.Data2View(data)
			if err == nil {
				viewTmpList = append(viewTmpList, view)
			}
		}
		viewList = viewTmpList
	}
	return
}
