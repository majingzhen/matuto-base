// Package convert
// @description <TODO description class purpose>
// @author
// @File: post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package view

import (
	"fmt"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
	"matuto-base/src/utils"
)

type PostViewUtils struct{}

func (viewUtils *PostViewUtils) Data2View(data *model.Post) (err error, view *PostView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("PostViewUtils View2Data error: %v", e)
			global.Logger.Error("PostViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp PostView

	tmp.Id = data.Id

	tmp.PostCode = data.PostCode

	tmp.PostName = data.PostName

	tmp.PostSort = data.PostSort

	tmp.Status = data.Status

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *PostViewUtils) View2Data(view *PostView) (err error, data *model.Post) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("PostViewUtils View2Data error: %v", e)
			global.Logger.Error("PostViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.Post

	tmp.Id = view.Id

	tmp.PostCode = view.PostCode

	tmp.PostName = view.PostName

	tmp.PostSort = view.PostSort

	tmp.Status = view.Status

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *PostViewUtils) View2DataList(viewList []*PostView) (err error, dataList []*model.Post) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("PostViewUtils View2DataList error: %v", e)
			global.Logger.Error("PostViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.Post
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

func (viewUtils *PostViewUtils) Data2ViewList(dataList []*model.Post) (err error, viewList []*PostView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("PostViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("PostViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*PostView
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

func (viewUtils *PostViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("PostViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("PostViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.Post); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
