// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_menu/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysMenuViewUtils struct{}

func (sysMenuViewUtils *SysMenuViewUtils) Data2View(data *model.SysMenu) (err error, view *SysMenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils View2Data error: %v", e)
			global.Logger.Error("SysMenuViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysMenuView

	tmp.Id = data.Id

	tmp.MenuName = data.MenuName

	tmp.ParentId = data.ParentId

	tmp.OrderNum = data.OrderNum

	tmp.Path = data.Path

	tmp.Component = data.Component

	tmp.Query = data.Query

	tmp.IsFrame = data.IsFrame

	tmp.IsCache = data.IsCache

	tmp.MenuType = data.MenuType

	tmp.Visible = data.Visible

	tmp.Status = data.Status

	tmp.Perms = data.Perms

	tmp.Icon = data.Icon

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (sysMenuViewUtils *SysMenuViewUtils) View2Data(view *SysMenuView) (err error, data *model.SysMenu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils View2Data error: %v", e)
			global.Logger.Error("SysMenuViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysMenu

	tmp.Id = view.Id

	tmp.MenuName = view.MenuName

	tmp.ParentId = view.ParentId

	tmp.OrderNum = view.OrderNum

	tmp.Path = view.Path

	tmp.Component = view.Component

	tmp.Query = view.Query

	tmp.IsFrame = view.IsFrame

	tmp.IsCache = view.IsCache

	tmp.MenuType = view.MenuType

	tmp.Visible = view.Visible

	tmp.Status = view.Status

	tmp.Perms = view.Perms

	tmp.Icon = view.Icon

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (sysMenuViewUtils *SysMenuViewUtils) View2DataList(viewList *[]SysMenuView) (err error, dataList *[]model.SysMenu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysMenuViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysMenu
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysMenuViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysMenuViewUtils *SysMenuViewUtils) Data2ViewList(dataList *[]model.SysMenu) (err error, viewList *[]SysMenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysMenuViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysMenuViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysMenuView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysMenuViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}
