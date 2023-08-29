// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/model"
	"manager-gin/src/common"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysDeptViewUtils struct{}

func (viewUtils *SysDeptViewUtils) Data2View(data *model.SysDept) (err error, view *SysDeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDeptViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysDeptView

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
func (viewUtils *SysDeptViewUtils) View2Data(view *SysDeptView) (err error, data *model.SysDept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils View2Data error: %v", e)
			global.Logger.Error("SysDeptViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysDept

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

func (viewUtils *SysDeptViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("SysDeptViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.SysDept); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err != nil {
				return err, nil
			} else {
				pageInfo.Rows = viewList
			}
		}
	}
	return nil, pageInfo
}

func (viewUtils *SysDeptViewUtils) Data2Tree(data *model.SysDept) (err error, view *SysDeptTreeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils Data2Tree error: %v", e)
			global.Logger.Error("SysDeptViewUtils.Data2Tree:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysDeptTreeView
	tmp.Id = data.Id
	tmp.Label = data.DeptName
	tmp.ParentId = data.ParentId
	view = &tmp
	return
}

func (viewUtils *SysDeptViewUtils) Data2TreeList(dataList []*model.SysDept) (err error, treeList []*SysDeptTreeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils Data2TreeList error: %v", e)
			global.Logger.Error("SysDeptViewUtils.Data2TreeList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var tmpList []*SysDeptTreeView
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

func (viewUtils *SysDeptViewUtils) View2DataList(viewList []*SysDeptView) (err error, dataList []*model.SysDept) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysDeptViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.SysDept
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

func (viewUtils *SysDeptViewUtils) Data2ViewList(dataList []*model.SysDept) (err error, viewList []*SysDeptView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysDeptViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysDeptViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*SysDeptView
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
