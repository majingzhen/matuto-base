// Package dao 自动生成模板 NoticeDao
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-14 15:32:04
package dao

import (
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
)

// NoticeDao 结构体
type NoticeDao struct{}

// Create 新增通知公告表记录
// Author matuto
func (dao *NoticeDao) Create(notice *model.Notice) (err error) {
	err = global.GormDao.Create(notice).Error
	return err
}

// DeleteByIds 批量删除通知公告表记录
// Author matuto
func (dao *NoticeDao) DeleteByIds(ids []string) (err error) {
	return global.GormDao.Delete(&[]model.Notice{}, "id in ?", ids).Error
}

// Update 更新通知公告表记录
// Author matuto
func (dao *NoticeDao) Update(notice *model.Notice) (err error) {
	err = global.GormDao.Updates(notice).Error
	return err
}

// Get 根据id获取通知公告表记录
// Author matuto
func (dao *NoticeDao) Get(id string) (err error, notice *model.Notice) {
	err = global.GormDao.Where("id = ?", id).First(notice).Error
	return
}

// Page 分页获取通知公告表记录
// Author matuto
func (dao *NoticeDao) Page(param *vo.NoticePageView) (err error, page *common.PageInfo) {
	db := global.GormDao.Model(&model.Notice{})

	if param.NoticeTitle != "" {
		db.Where("notice_title like ?", "%"+param.NoticeTitle+"%")
	}

	if param.NoticeType != "" {
		db.Where("notice_type = ?", param.NoticeType)
	}

	if param.Status != "" {
		db.Where("status = ?", param.Status)
	}

	if param.CreateTime != "" {
		db.Where("create_time = ?", param.CreateTime)
	}

	if param.BeginUpdateTime != "" && param.EndUpdateTime != "" {
		db.Where("update_time between ? and ?", param.BeginUpdateTime, param.EndUpdateTime)
	}

	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.Notice
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

// List 获取通知公告表记录
// Author matuto
func (dao *NoticeDao) List(param *vo.NoticeQueryView) (err error, dataList []*model.Notice) {
	db := global.GormDao.Model(&model.Notice{})

	if param.NoticeTitle != "" {
		db.Where("notice_title like ?", "%"+param.NoticeTitle+"%")
	}

	if param.NoticeType != "" {
		db.Where("notice_type = ?", param.NoticeType)
	}

	if param.Status != "" {
		db.Where("status = ?", param.Status)
	}

	if param.CreateTime != "" {
		db.Where("create_time = ?", param.CreateTime)
	}

	if param.BeginUpdateTime != "" && param.EndUpdateTime != "" {
		db.Where("update_time between ? and ?", param.BeginUpdateTime, param.EndUpdateTime)
	}

	db.Order("create_time desc")
	err = db.Find(&dataList).Error
	return err, dataList
}
