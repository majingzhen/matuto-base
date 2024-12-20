// Package convert
// @description <TODO description class purpose>
// @author
// @File: user
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

type UserViewUtils struct{}

func (viewUtils *UserViewUtils) Data2View(data *model.User) (err error, view *UserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("UserViewUtils View2Data error: %v", e)
			global.Logger.Error("UserViewUtils.Data2View:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp UserView

	tmp.Id = data.Id

	tmp.DeptId = data.DeptId

	tmp.UserName = data.UserName

	tmp.NickName = data.NickName

	tmp.UserType = data.UserType

	tmp.Email = data.Email

	tmp.PhoneNumber = data.PhoneNumber

	tmp.Sex = data.Sex

	tmp.Avatar = data.Avatar

	tmp.Password = data.Password

	tmp.Salt = data.Salt

	tmp.Status = data.Status

	tmp.LoginIp = data.LoginIp

	tmp.LoginDate = utils.Time2Str(data.LoginDate)

	tmp.CreateBy = data.CreateBy

	tmp.CreateTime = utils.Time2Str(data.CreateTime)

	tmp.UpdateBy = data.UpdateBy

	tmp.UpdateTime = utils.Time2Str(data.UpdateTime)

	tmp.Remark = data.Remark

	view = &tmp
	return
}
func (viewUtils *UserViewUtils) View2Data(view *UserView) (err error, data *model.User) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("UserViewUtils View2Data error: %v", e)
			global.Logger.Error("UserViewUtils.View2Data:格式转换异常",
				zap.Any("error", e))
		}
	}()
	var tmp model.User

	tmp.Id = view.Id

	tmp.DeptId = view.DeptId

	tmp.UserName = view.UserName

	tmp.NickName = view.NickName

	tmp.UserType = view.UserType

	tmp.Email = view.Email

	tmp.PhoneNumber = view.PhoneNumber

	tmp.Sex = view.Sex

	tmp.Avatar = view.Avatar

	tmp.Password = view.Password

	tmp.Salt = view.Salt

	tmp.Status = view.Status

	tmp.LoginIp = view.LoginIp

	tmp.LoginDate = utils.Str2Time(view.LoginDate)

	tmp.CreateBy = view.CreateBy

	tmp.CreateTime = utils.Str2Time(view.CreateTime)

	tmp.UpdateBy = view.UpdateBy

	tmp.UpdateTime = utils.Str2Time(view.UpdateTime)

	tmp.Remark = view.Remark

	data = &tmp
	return
}

func (viewUtils *UserViewUtils) View2DataList(viewList []*UserView) (err error, dataList []*model.User) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("UserViewUtils View2DataList error: %v", e)
			global.Logger.Error("UserViewUtils.View2DataList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if viewList != nil {
		var dataTmpList []*model.User
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

func (viewUtils *UserViewUtils) Data2ViewList(dataList []*model.User) (err error, viewList []*UserView) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("UserViewUtils Data2ViewList error: %v", e)
			global.Logger.Error("UserViewUtils.Data2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if dataList != nil {
		var viewTmpList []*UserView
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

func (viewUtils *UserViewUtils) PageData2ViewList(pageInfo *common.PageInfo) (err error, res *common.PageInfo) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("UserViewUtils PageData2ViewList error: %v", e)
			global.Logger.Error("UserViewUtils.PageData2ViewList:格式转换异常",
				zap.Any("error", e))
		}
	}()
	if pageInfo != nil && pageInfo.Rows != nil {
		if p, ok := pageInfo.Rows.([]*model.User); ok {
			if err, viewList := viewUtils.Data2ViewList(p); err == nil {
				pageInfo.Rows = viewList
			}
		}
	}
	res = pageInfo
	return
}
