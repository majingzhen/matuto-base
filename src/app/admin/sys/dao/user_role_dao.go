// Package dao 自动生成模板 UserRoleDao
// @description <TODO description class purpose>
// @author
// @File: user_role
// @version 1.0.0
// @create 2023-08-21 17:37:56
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/global"
)

// UserRoleDao 结构体

type UserRoleDao struct{}

// CreateBatch 批量创建UserRole记录
func (dao *UserRoleDao) CreateBatch(tx *gorm.DB, roles []model.UserRole) error {
	return tx.Create(&roles).Error
}

// DeleteByUserIds 根据用户id删除用户角色关联数据
func (dao *UserRoleDao) DeleteByUserIds(tx *gorm.DB, ids []string) error {
	return tx.Delete(&[]model.UserRole{}, "user_id in ?", ids).Error
}

// CountUserRoleByRoleId 根据角色id查询用户数量
func (dao *UserRoleDao) CountUserRoleByRoleId(id string) (error, int64) {
	var total int64
	err := global.GormDao.Model(&model.UserRole{}).Where("role_id = ?", id).Count(&total).Error
	return err, total

}

// DeleteUserRoleInfo 根据用户id和角色id删除用户角色关联数据
func (dao *UserRoleDao) DeleteUserRoleInfo(userId, roleId string) error {
	return global.GormDao.Delete(&model.UserRole{}, "user_id = ? and role_id = ?", userId, roleId).Error
}

// DeleteUsersRoleInfo 根据用户id集合和角色id删除用户角色关联数据
func (dao *UserRoleDao) DeleteUsersRoleInfo(roleId string, userIds []string) error {
	return global.GormDao.Delete(&model.UserRole{}, "user_id in ? and role_id = ?", userIds, roleId).Error
}

// InsertUsersRoleInfo 批量插入用户角色关联数据
func (dao *UserRoleDao) InsertUsersRoleInfo(roleId string, userIds []string) error {
	var userRoles []model.UserRole
	for _, userId := range userIds {
		userRoles = append(userRoles, model.UserRole{RoleId: roleId, UserId: userId})
	}
	return global.GormDao.Create(&userRoles).Error
}
