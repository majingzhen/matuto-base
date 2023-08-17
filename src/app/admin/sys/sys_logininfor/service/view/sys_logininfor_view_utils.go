// Package view
// @description <TODO description class purpose>
// @author
// @File: sys_logininfor
// @version 1.0.0
// @create 2023-08-08 10:06:19
package view

import (
	"fmt"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/sys/sys_logininfor/model"
	"manager-gin/src/global"
	"manager-gin/src/utils"
)

type SysLogininforViewUtils struct{}

func (sysLogininforViewUtils *SysLogininforViewUtils) Data2View(data *model.SysLogininfor) (err error, view *SysLogininforView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysLogininforViewUtils View2Data error: %v", e)
			global.Logger.Error("SysLogininforViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp SysLogininforView

	tmp.Browser = data.Browser

	tmp.Id = data.Id

	tmp.Ipaddr = data.Ipaddr

	tmp.LoginLocation = data.LoginLocation

	tmp.LoginTime = utils.Time2Str(data.LoginTime)

	tmp.Msg = data.Msg

	tmp.Os = data.Os

	tmp.Status = data.Status

	tmp.UserName = data.UserName

	view = &tmp
	return
}
func (sysLogininforViewUtils *SysLogininforViewUtils) View2Data(view *SysLogininforView) (err error, data *model.SysLogininfor) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysLogininforViewUtils View2Data error: %v", e)
			global.Logger.Error("SysLogininforViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.SysLogininfor

	tmp.Browser = view.Browser

	tmp.Id = view.Id

	tmp.Ipaddr = view.Ipaddr

	tmp.LoginLocation = view.LoginLocation

	tmp.LoginTime = utils.Str2Time(view.LoginTime)

	tmp.Msg = view.Msg

	tmp.Os = view.Os

	tmp.Status = view.Status

	tmp.UserName = view.UserName

	data = &tmp
	return
}

func (sysLogininforViewUtils *SysLogininforViewUtils) View2DataList(viewList *[]SysLogininforView) (err error, dataList *[]model.SysLogininfor) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysLogininforViewUtils View2DataList error: %v", e)
			global.Logger.Error("SysLogininforViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []model.SysLogininfor
		for i := range *dataList {
			view := (*viewList)[i]
			err, data := sysLogininforViewUtils.View2Data(&view)
			if err == nil {
				dataTmpList = append(dataTmpList, *data)
			}
		}
		dataList = &dataTmpList
	}
	return
}

func (sysLogininforViewUtils *SysLogininforViewUtils) Data2ViewList(dataList *[]model.SysLogininfor) (err error, viewList *[]SysLogininforView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("SysLogininforViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("SysLogininforViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []SysLogininforView
		for i := range *dataList {
			data := (*dataList)[i]
			err, view := sysLogininforViewUtils.Data2View(&data)
			if err == nil {
				viewTmpList = append(viewTmpList, *view)
			}
		}
		viewList = &viewTmpList
	}
	return
}