// Package dao 自动生成模板 ConfigDao
// @description <TODO description class purpose>
// @author
// @File: config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package dao

import (
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/app/admin/sys/service/config/view"
	"matuto-base/src/common"
	"matuto-base/src/global"
)

// ConfigDao 结构体
type ConfigDao struct{}

// Create 创建Config记录
// Author
func (dao *ConfigDao) Create(sysConfig model.Config) (err error) {
	err = global.GormDao.Create(&sysConfig).Error
	return err
}

// DeleteByIds 批量删除Config记录
// Author
func (dao *ConfigDao) DeleteByIds(ids []string) (err error) {
	err = global.GormDao.Delete(&[]model.Config{}, "id in ?", ids).Error
	return err
}

// Update 更新Config记录
// Author
func (dao *ConfigDao) Update(sysConfig model.Config) (err error) {
	err = global.GormDao.Updates(&sysConfig).Error
	return err
}

// Get 根据id获取Config记录
// Author
func (dao *ConfigDao) Get(id string) (err error, sysConfig *model.Config) {
	err = global.GormDao.Where("id = ?", id).First(&sysConfig).Error
	return
}

// Page 分页获取Config记录
// Author
func (dao *ConfigDao) Page(param *view.ConfigPageView) (err error, res *common.PageInfo) {
	// 创建db
	db := global.GormDao.Model(&model.Config{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.ConfigName != "" {
		db.Where("config_name like ?", "%"+param.ConfigName+"%")
	}
	if param.ConfigKey != "" {
		db.Where("config_key like ?", "%"+param.ConfigKey+"%")
	}
	if param.ConfigType != "" {
		db.Where("config_type = ?", param.ConfigType)
	}
	page := common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}

	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.Config
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

// List 获取Config记录
// Author
func (dao *ConfigDao) List(data *model.Config) (err error, datas []*model.Config) {
	var rows []*model.Config
	db := global.GormDao.Model(&model.Config{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	db.Order("create_time desc")
	err = db.Find(&rows).Error
	datas = rows
	return err, datas
}

// SelectConfigByKey 根据key查询Config记录
func (dao *ConfigDao) SelectConfigByKey(key string) (error, *model.Config) {
	var rows []*model.Config
	db := global.GormDao.Model(&model.Config{})
	db.Where("config_key = ?", key)
	err := db.Find(&rows).Error
	if rows != nil && len(rows) > 0 {
		return err, rows[0]
	}
	return err, nil
}
