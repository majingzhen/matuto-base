// Package utils 提供通用的VO和DO对象转换工具类
package convert

import (
	"fmt"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"matuto-base/src/common"
	"matuto-base/src/global"
)

// Data2View 单个DO转VO
func Data2View[V any, D any](data *D) (err error, view *V) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("Data2View error: %v", e)
			global.Logger.Error("Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()

	if data == nil {
		return nil, nil
	}

	var v V
	view = &v

	if err := copier.Copy(view, data); err != nil {
		return err, nil
	}

	return nil, view
}

// View2Data 单个VO转DO
func View2Data[V any, D any](view *V) (err error, data *D) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("View2Data error: %v", e)
			global.Logger.Error("View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()

	if view == nil {
		return nil, nil
	}

	var d D
	data = &d

	if err := copier.Copy(data, view); err != nil {
		return err, nil
	}

	return nil, data
}

// Data2ViewList DO列表转VO列表
func Data2ViewList[V any, D any](dataList []*D) (err error, viewList []*V) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("Data2ViewList error: %v", e)
			global.Logger.Error("Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()

	if dataList == nil {
		return nil, nil
	}

	viewList = make([]*V, len(dataList))
	if err := copier.Copy(&viewList, dataList); err != nil {
		return err, nil
	}

	return nil, viewList
}

// View2DataList VO列表转DO列表
func View2DataList[V any, D any](viewList []*V) (err error, dataList []*D) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("View2DataList error: %v", e)
			global.Logger.Error("View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()

	if viewList == nil {
		return nil, nil
	}

	dataList = make([]*D, len(viewList))
	if err := copier.Copy(&dataList, viewList); err != nil {
		return err, nil
	}

	return nil, dataList
}

// PageData2ViewList 分页数据DO列表转VO列表
func PageData2ViewList[V any, D any](pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("PageData2ViewList error: %v", e)
			global.Logger.Error("PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()

	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*D); ok {
			if err, viewList := Data2ViewList[V, D](p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
