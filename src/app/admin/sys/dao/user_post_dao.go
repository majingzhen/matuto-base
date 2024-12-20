// Package dao 自动生成模板 UserPostDao
// @description <TODO description class purpose>
// @author
// @File: user_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/global"
)

// UserPostDao 结构体

type UserPostDao struct{}

// CreateBatch 批量创建UserPost记录
func (dao *UserPostDao) CreateBatch(tx *gorm.DB, posts []model.UserPost) error {
	return global.GormDao.Create(&posts).Error
}

// DeleteByUserIds 根据用户id批量删除UserPost记录
func (dao *UserPostDao) DeleteByUserIds(tx *gorm.DB, ids []string) error {
	return global.GormDao.Delete(&model.UserPost{}, "user_id in ?", ids).Error
}
