// Package convert
// @description <TODO description class purpose>
// @author
// @File: sys_menu
// @version 1.0.0
// @create 2023-08-20 21:21:34
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type MenuViewUtils struct{}

func (viewUtils *MenuViewUtils) Data2View(data *model.Menu) (err error, view *MenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("MenuViewUtils View2Data error: %v", e)
			global.Logger.Error("MenuViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp MenuView

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
func (viewUtils *MenuViewUtils) View2Data(view *MenuView) (err error, data *model.Menu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("MenuViewUtils View2Data error: %v", e)
			global.Logger.Error("MenuViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.Menu

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

func (viewUtils *MenuViewUtils) View2DataList(viewList []*MenuView) (err error, dataList []*model.Menu) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("MenuViewUtils View2DataList error: %v", e)
			global.Logger.Error("MenuViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.Menu
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

func (viewUtils *MenuViewUtils) Data2ViewList(dataList []*model.Menu) (err error, viewList []*MenuView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("MenuViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("MenuViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*MenuView
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
