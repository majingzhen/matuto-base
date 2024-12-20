// Package convert
// @description <TODO description class purpose>
// @author
// @File: dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type DeptViewUtils struct{}

func (viewUtils *DeptViewUtils) Data2View(data *model.Dept) (err error, view *DeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DeptViewUtils View2Data error: %v", e)
			global.Logger.Error("DeptViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp DeptView

	tmp.Id = data.Id

	tmp.ParentId = data.ParentId

	tmp.Ancestors = data.Ancestors

	tmp.DeptName = data.DeptName

	tmp.OrderNum = data.OrderNum

	tmp.Leader = data.Leader

	tmp.Phone = data.Phone

	tmp.Email = data.Email

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	view = &tmp
	return
}
func (viewUtils *DeptViewUtils) View2Data(view *DeptView) (err error, data *model.Dept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DeptViewUtils View2Data error: %v", e)
			global.Logger.Error("DeptViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.Dept

	tmp.Id = view.Id

	tmp.ParentId = view.ParentId

	tmp.Ancestors = view.Ancestors

	tmp.DeptName = view.DeptName

	tmp.OrderNum = view.OrderNum

	tmp.Leader = view.Leader

	tmp.Phone = view.Phone

	tmp.Email = view.Email

	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	data = &tmp
	return
}

func (viewUtils *DeptViewUtils) Data2Tree(data *model.Dept) (err error, view *DeptTreeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DeptViewUtils Data2Tree error: %v", e)
			global.Logger.Error("DeptViewUtils.Data2Tree:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp DeptTreeView
	tmp.Id = data.Id
	tmp.Label = data.DeptName
	tmp.ParentId = data.ParentId
	view = &tmp
	return
}

func (viewUtils *DeptViewUtils) Data2TreeList(dataList []*model.Dept) (err error, treeList []*DeptTreeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DeptViewUtils Data2TreeList error: %v", e)
			global.Logger.Error("DeptViewUtils.Data2TreeList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var tmpList []*DeptTreeView
		for i := range dataList {
			data := (dataList)[i]
			err, tree := viewUtils.Data2Tree(data)
			if err == nil {
				tmpList = append(tmpList, tree)
			}
		}
		treeList = tmpList
	}
	return
}

func (viewUtils *DeptViewUtils) View2DataList(viewList []*DeptView) (err error, dataList []*model.Dept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DeptViewUtils View2DataList error: %v", e)
			global.Logger.Error("DeptViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.Dept
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

func (viewUtils *DeptViewUtils) Data2ViewList(dataList []*model.Dept) (err error, viewList []*DeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("DeptViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("DeptViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*DeptView
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
