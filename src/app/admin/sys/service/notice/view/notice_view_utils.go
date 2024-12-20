// Package convert 自动生成模板 NoticeViewUtils
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-12 13:40:27
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type NoticeViewUtils struct{}

// Data2View Do转Vo
func (viewUtils *NoticeViewUtils) Data2View(data *model.Notice) (err error, view *NoticeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("NoticeViewUtils View2Data error: %v", e)
			global.Logger.Error("NoticeViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	view = &NoticeView{}
	view.Id = data.Id
	view.NoticeTitle = data.NoticeTitle
	view.NoticeType = data.NoticeType
	view.NoticeContent = string(data.NoticeContent)
	view.Status = data.Status
	view.CreateBy = data.CreateBy
	view.CreateTime = utils.Time2Str(data.CreateTime)
	view.UpdateBy = data.UpdateBy
	view.UpdateTime = utils.Time2Str(data.UpdateTime)
	view.Remark = data.Remark

	return
}

// View2Data Vo转Do
func (viewUtils *NoticeViewUtils) View2Data(view *NoticeView) (err error, data *model.Notice) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("NoticeViewUtils View2Data error: %v", e)
			global.Logger.Error("NoticeViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	data = &model.Notice{}
	data.Id = view.Id
	data.NoticeTitle = view.NoticeTitle
	data.NoticeType = view.NoticeType
	data.NoticeContent = []byte(view.NoticeContent)
	data.Status = view.Status
	data.CreateBy = view.CreateBy
	data.CreateTime = utils.Str2Time(view.CreateTime)
	data.UpdateBy = view.UpdateBy
	data.UpdateTime = utils.Str2Time(view.UpdateTime)
	data.Remark = view.Remark

	return
}

// View2DataList VoList转DoList
func (viewUtils *NoticeViewUtils) View2DataList(viewList []*NoticeView) (err error, dataList []*model.Notice) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("NoticeViewUtils View2DataList error: %v", e)
			global.Logger.Error("NoticeViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.Notice
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

// Data2ViewList DoList转VoList
func (viewUtils *NoticeViewUtils) Data2ViewList(dataList []*model.Notice) (err error, viewList []*NoticeView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("NoticeViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("NoticeViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*NoticeView
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

// PageData2ViewList 分页数据转换
func (viewUtils *NoticeViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("NoticeViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("NoticeViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.Notice); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
