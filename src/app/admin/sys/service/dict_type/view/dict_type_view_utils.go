// Package convert
// @description <TODO description class purpose>
// @author
// @File: dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type DictTypeViewUtils struct{}

func (viewUtils *DictTypeViewUtils) Data2View(data *model.DictType) (err error, view *DictTypeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictTypeViewUtils View2Data error: %v", e)
			global.Logger.Error("DictTypeViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp DictTypeView
	tmp.Id = data.Id
	tmp.DictName = data.DictName
	tmp.DictType = data.DictType
	tmp.Status = data.Status
	tmp.CreateBy = data.CreateBy
	tmp.CreateTime = utils.Time2Str(data.CreateTime)
	tmp.UpdateBy = data.UpdateBy
	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)
	tmp.Remark = data.Remark
	view = &tmp
	return
}

func (viewUtils *DictTypeViewUtils) View2Data(view *DictTypeView) (err error, data *model.DictType) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictTypeViewUtils View2Data error: %v", e)
			global.Logger.Error("DictTypeViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.DictType
	tmp.Id = view.Id
	tmp.DictName = view.DictName
	tmp.DictType = view.DictType
	tmp.Status = view.Status
	tmp.CreateBy = view.CreateBy
	tmp.CreateTime = utils.Str2Time(view.CreateTime)
	tmp.UpdateBy = view.UpdateBy
	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)
	tmp.Remark = view.Remark
	data = &tmp
	return
}

func (viewUtils *DictTypeViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictTypeViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("DictTypeViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.DictType); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}

func (viewUtils *DictTypeViewUtils) View2DataList(viewList []*DictTypeView) (err error, dataList []*model.DictType) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictTypeViewUtils View2DataList error: %v", e)
			global.Logger.Error("DictTypeViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.DictType
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

func (viewUtils *DictTypeViewUtils) Data2ViewList(dataList []*model.DictType) (err error, viewList []*DictTypeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DictTypeViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("DictTypeViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*DictTypeView
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
