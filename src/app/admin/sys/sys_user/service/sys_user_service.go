// Package service 自动生成模板 SysUserService
// @description <TODO description class purpose>
// @author
// @File: sys_user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package service

import (
	"manager-gin/src/app/admin/sys/sys_user/model"
	"manager-gin/src/app/admin/sys/sys_user/service/view"
	"manager-gin/src/common"
)

var sysUserDao = model.SysUserDaoApp
var viewUtils = view.SysUserViewUtilsApp

type SysUserService struct{}

// Create 创建SysUser记录
// Author
func (service *SysUserService) Create(sysUserView *view.SysUserView) (err error) {
	err1, sysUser := viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	err2 := sysUserDao.Create(*sysUser)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除SysUser记录
// Author
func (service *SysUserService) Delete(id string) (err error) {
	err = sysUserDao.Delete(id)
	return err
}

// DeleteByIds 批量删除SysUser记录
// Author
func (service *SysUserService) DeleteByIds(ids []string) (err error) {
	err = sysUserDao.DeleteByIds(ids)
	return err
}

// Update 更新SysUser记录
// Author
func (service *SysUserService) Update(id string, sysUserView *view.SysUserView) (err error) {
	sysUserView.Id = id
	err1, sysUser := viewUtils.View2Data(sysUserView)
	if err1 != nil {
		return err1
	}
	err = sysUserDao.Update(*sysUser)
	return err
}

// Get 根据id获取SysUser记录
// Author
func (service *SysUserService) Get(id string) (err error, sysUserView *view.SysUserView) {
	err1, sysUser := sysUserDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := viewUtils.Data2View(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}

// List 分页获取SysUser记录
// Author
func (service *SysUserService) List(info *common.PageInfo) (err error) {
	err1, sysUsers, total := sysUserDao.List(info)
	if err1 != nil {
		return err1
	}
	info.Total = total
	err2, viewList := viewUtils.Data2ViewList(sysUsers)
	if err2 != nil {
		return err2
	}
	info.Rows = viewList
	return err
}

// GetByUserName 根据userName获取SysUser记录
// Author
func (service *SysUserService) GetByUserName(userName string) (err error, sysUserView *view.SysUserView) {
	err1, sysUser := sysUserDao.GetByUserName(userName)
	if err1 != nil {
		return err1, nil
	}
	err2, sysUserView := viewUtils.Data2View(sysUser)
	if err2 != nil {
		return err2, nil
	}
	return
}

// IsAdmin 用户是否管理员
func (service *SysUserService) IsAdmin(userId string) (itIs bool) {
	if common.SYSTEM_ADMIN_ID == userId {
		itIs = true
	}
	return
}
