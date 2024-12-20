// Package dao 自动生成模板 RoleDeptDao
// @description <TODO description class purpose>
// @author
// @File: role_dept
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/model"
)

// RoleDeptDao 结构体

type RoleDeptDao struct{}

// DeleteRoleDeptByRoleIds 根据角色id集合删除角色部门关联数据
func (dao *RoleDeptDao) DeleteRoleDeptByRoleIds(tx *gorm.DB, ids []string) error {
	return tx.Delete(&[]model.RoleDept{}, "role_id in ?", ids).Error
}

// DeleteRoleDeptByRoleId 根据角色id集合删除角色部门关联数据
func (dao *RoleDeptDao) DeleteRoleDeptByRoleId(tx *gorm.DB, id string) error {
	return tx.Delete(&[]model.RoleDept{}, "role_id = ?", id).Error

}

// CreateBatch 批量创建RoleDept记录
func (dao *RoleDeptDao) CreateBatch(tx *gorm.DB, depts []model.RoleDept) error {
	return tx.Create(&depts).Error
}
