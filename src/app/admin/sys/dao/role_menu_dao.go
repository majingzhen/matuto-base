// Package dao 自动生成模板 RoleMenuDao
// @description <TODO description class purpose>
// @author
// @File: role_menu
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/model"
)

// RoleMenuDao 结构体

type RoleMenuDao struct{}

// DeleteRoleMenuByRoleId 根据角色id删除角色菜单关联数据
func (dao *RoleMenuDao) DeleteRoleMenuByRoleId(tx *gorm.DB, id string) error {
	return tx.Delete(&[]model.RoleMenu{}, "role_id = ?", id).Error
}

// CreateBatch 批量插入
func (dao *RoleMenuDao) CreateBatch(tx *gorm.DB, menus []model.RoleMenu) error {
	return tx.Create(&menus).Error
}

// DeleteRoleMenuByRoleIds 根据角色id集合删除角色菜单关联数据
func (dao *RoleMenuDao) DeleteRoleMenuByRoleIds(tx *gorm.DB, ids []string) error {
	return tx.Delete(&[]model.RoleMenu{}, "role_id in ?", ids).Error
}
