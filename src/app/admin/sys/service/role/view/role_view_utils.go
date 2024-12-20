// Package convert
// @description <TODO description class purpose>
// @author
// @File: role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package view

import (
	"fmt"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
	"matuto-base/src/utils"

	"go.uber.org/zap"
)

type RoleViewUtils struct{}

func (viewUtils *RoleViewUtils) Data2View(data *model.Role) (err error, view *RoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("RoleViewUtils View2Data error: %v", e)
			global.Logger.Error("RoleViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp RoleView

	tmp.Id = data.Id

	tmp.RoleName = data.RoleName

	tmp.RoleKey = data.RoleKey

	tmp.RoleSort = data.RoleSort

	tmp.DataScope = data.DataScope

	if data.MenuCheckStrictly == 1 {
		tmp.MenuCheckStrictly = true
	} else {
		tmp.MenuCheckStrictly = false
	}
	if data.DeptCheckStrictly == 1 {
		tmp.DeptCheckStrictly = true
	} else {
		tmp.DeptCheckStrictly = false
	}

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *RoleViewUtils) View2Data(view *RoleView) (err error, data *model.Role) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("RoleViewUtils View2Data error: %v", e)
			global.Logger.Error("RoleViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.Role

	tmp.Id = view.Id

	tmp.RoleName = view.RoleName

	tmp.RoleKey = view.RoleKey

	tmp.RoleSort = view.RoleSort

	tmp.DataScope = view.DataScope

	if view.MenuCheckStrictly {
		tmp.MenuCheckStrictly = 1
	} else {
		tmp.MenuCheckStrictly = 0
	}

	if view.DeptCheckStrictly {
		tmp.DeptCheckStrictly = 1
	} else {
		tmp.DeptCheckStrictly = 0
	}
	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *RoleViewUtils) View2DataList(viewList []*RoleView) (err error, dataList []*model.Role) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("RoleViewUtils View2DataList error: %v", e)
			global.Logger.Error("RoleViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.Role
		for i := range viewList {
			view := (viewList)[i]
			err, data := viewUtils.View2Data(view)
			if err == nil {
				dataTmpList = append(dataTmpList, data)
			}
		}
		dataList = dataTmpList
	}
	return
}

func (viewUtils *RoleViewUtils) Data2ViewList(dataList []*model.Role) (err error, viewList []*RoleView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("RoleViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("RoleViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*RoleView
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

func (viewUtils *RoleViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("RoleViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("RoleViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.Role); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
