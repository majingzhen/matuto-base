// Package dao 自动生成模板 DictTypeDao
// @description <TODO description class purpose>
// @author
// @File: dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package dao

import (
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/app/admin/sys/service/dict_type/view"
	"matuto-base/src/common"
	"matuto-base/src/global"
)

// DictTypeDao 结构体

type DictTypeDao struct{}

// Create 创建DictType记录
// Author
func (dao *DictTypeDao) Create(sysDictType model.DictType) (err error) {
	err = global.GormDao.Create(&sysDictType).Error
	return err
}

// Delete 删除DictType记录
// Author
func (dao *DictTypeDao) Delete(id string) (err error) {
	err = global.GormDao.Delete(&[]model.DictType{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除DictType记录
// Author
func (dao *DictTypeDao) DeleteByIds(ids []string) (err error) {
	err = global.GormDao.Delete(&[]model.DictType{}, "id in ?", ids).Error
	return err
}

// Update 更新DictType记录
// Author
func (dao *DictTypeDao) Update(sysDictType model.DictType) (err error) {
	err = global.GormDao.Updates(&sysDictType).Error
	return err
}

// Get 根据id获取DictType记录
// Author
func (dao *DictTypeDao) Get(id string) (err error, sysDictType *model.DictType) {
	err = global.GormDao.Where("id = ?", id).First(&sysDictType).Error
	return
}

// Page 分页获取DictType记录
// Author
func (dao *DictTypeDao) Page(param *view.DictTypePageView) (err error, page *common.PageInfo) {
	// 创建db
	db := global.GormDao.Model(&model.DictType{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.DictName != "" {
		db = db.Where("dict_name like ?", "%"+param.DictName+"%")
	}
	if param.DictType != "" {
		db = db.Where("dict_type like ?", "%"+param.DictType+"%")
	}
	if param.Status != "" {
		db = db.Where("status = ?", param.Status)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	err = db.Count(&page.Total).Error
	if err != nil {
		return
	}
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.DictType
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

func (dao *DictTypeDao) SelectDictTypeAll() (err error, datas []*model.DictType) {
	err = global.GormDao.Find(&datas).Error
	return
}
