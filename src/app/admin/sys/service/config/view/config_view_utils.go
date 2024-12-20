// Package convert
// @description <TODO description class purpose>
// @author
// @File: config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type ConfigViewUtils struct{}

func (viewUtils *ConfigViewUtils) Data2View(data *model.Config) (err error, view *ConfigView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("ConfigViewUtils View2Data error: %v", e)
			global.Logger.Error("ConfigViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp ConfigView

	tmp.Id = data.Id

	tmp.ConfigName = data.ConfigName

	tmp.ConfigKey = data.ConfigKey

	tmp.ConfigValue = data.ConfigValue

	tmp.ConfigType = data.ConfigType

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *ConfigViewUtils) View2Data(view *ConfigView) (err error, data *model.Config) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("ConfigViewUtils View2Data error: %v", e)
			global.Logger.Error("ConfigViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.Config

	tmp.Id = view.Id

	tmp.ConfigName = view.ConfigName

	tmp.ConfigKey = view.ConfigKey

	tmp.ConfigValue = view.ConfigValue

	tmp.ConfigType = view.ConfigType

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *ConfigViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("ConfigViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("ConfigViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.Config); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}

func (viewUtils *ConfigViewUtils) View2DataList(viewList []*ConfigView) (err error, dataList []*model.Config) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("ConfigViewUtils View2DataList error: %v", e)
			global.Logger.Error("ConfigViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.Config
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

func (viewUtils *ConfigViewUtils) Data2ViewList(dataList []*model.Config) (err error, viewList []*ConfigView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("ConfigViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("ConfigViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*ConfigView
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
