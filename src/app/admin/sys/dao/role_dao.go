// Package dao 自动生成模板 RoleDao
// @description <TODO description class purpose>
// @author
// @File: role
// @version 1.0.0
// @create 2023-08-18 14:00:53
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/common/constants"
	"matuto-base/src/global"
)

// RoleDao 结构体

type RoleDao struct{}

// Create 创建Role记录
// Author
func (dao *RoleDao) Create(tx *gorm.DB, sysRole model.Role) (err error) {
	err = tx.Create(&sysRole).Error
	return err
}

// Delete 删除Role记录
// Author
func (dao *RoleDao) Delete(id string) (err error) {
	err = global.GormDao.Delete(&[]model.Role{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除Role记录
// Author
func (dao *RoleDao) DeleteByIds(tx *gorm.DB, ids []string) (err error) {
	err = tx.Delete(&[]model.Role{}, "id in ?", ids).Error
	return err
}

// Update 更新Role记录
// Author
func (dao *RoleDao) Update(tx *gorm.DB, sysRole *model.Role) (err error) {
	err = tx.Updates(sysRole).Error
	return err
}

// Get 根据id获取Role记录
// Author
func (dao *RoleDao) Get(id string) (err error, sysRole *model.Role) {
	err = global.GormDao.Where("id = ?", id).First(&sysRole).Error
	return
}

// Page 分页获取Role记录
// Author
func (dao *RoleDao) Page(param *vo.RolePageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GormDao.Table("sys_role r")
	db.Select("distinct r.id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly,r.status, r.create_time, r.remark ")
	db.Joins("left join sys_user_role ur on ur.role_id = r.id")
	db.Joins("left join sys_user u on u.id = ur.user_id")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.Id != "" {
		db.Where("ID = ?", param.Id)
	}
	if param.RoleName != "" {
		db.Where("role_name = ?", "%"+param.RoleName+"%")
	}
	if param.RoleKey != "" {
		db.Where("role_key = ?", "%"+param.RoleKey+"%")
	}
	if param.Status != "" {
		db.Where("status = ?", param.Status)
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var tmp []*model.Role
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	page.Rows = tmp
	return err, page
}

// List 获取Role记录
// Author
func (dao *RoleDao) List(data *model.Role) (err error, datas []*model.Role) {
	var rows []*model.Role
	model := global.GormDao.Table("sys_role r")
	model.Select("distinct r.id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly,r.status, r.create_time, r.remark ")
	model.Joins("left join sys_user_role ur on ur.role_id = r.id")
	model.Joins("left join sys_user u on u.id = ur.user_id")
	model.Joins("left join sys_dept d on u.dept_id = d.id")
	if data.Id != "" {
		model = model.Where("ID = ?", data.Id)
	}
	if data.RoleName != "" {
		model = model.Where("role_name = ?", "%"+data.RoleName+"%")
	}
	if data.RoleKey != "" {
		model = model.Where("role_key = ?", "%"+data.RoleKey+"%")
	}
	if data.Status != "" {
		model = model.Where("status = ?", data.Status)
	}
	if data.DataScopeSql != "" {
		model = model.Where(data.DataScopeSql)
	}
	model.Order("create_time desc")
	err = model.Find(&rows).Error
	datas = rows
	return err, datas
}

// GetRoleByUserId 根据用户获取角色集合
func (dao *RoleDao) GetRoleByUserId(userId string) (err error, roles []*model.Role) {
	var tmp []*model.Role
	model := global.GormDao.Table("sys_role r")
	model.Select("distinct r.id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly,r.status, r.create_time, r.remark ")
	model.Joins("left join sys_user_role ur on ur.role_id = r.id")
	model.Where("ur.user_id = ? and r.status = ?", userId, constants.STATUS_NORMAL)
	err = model.Find(&tmp).Error
	roles = tmp
	return err, roles
}

// CheckRoleNameUnique 校验角色名称是否唯一
func (dao *RoleDao) CheckRoleNameUnique(name string) (error, *model.Role) {
	var data []*model.Role
	if err := global.GormDao.Table("sys_role").Where("role_name = ?", name).Find(&data).Error; err != nil {
		return err, nil
	} else {
		if data != nil && len(data) > 0 {
			return nil, data[0]
		} else {
			return nil, nil
		}
	}
}

// CheckRoleKeyUnique 校验角色权限是否唯一
func (dao *RoleDao) CheckRoleKeyUnique(key string) (error, *model.Role) {
	var data []*model.Role
	if err := global.GormDao.Table("sys_role").Where("role_key = ?", key).Find(&data).Error; err != nil {
		return err, nil
	} else {
		if data != nil && len(data) > 0 {
			return nil, data[0]
		} else {
			return nil, nil
		}
	}
}
