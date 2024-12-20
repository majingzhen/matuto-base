// Package dao 自动生成模板 TableDao
// @description <TODO description class purpose>
// @author
// @File: table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/gen/api/vo"
	"matuto-base/src/app/admin/gen/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
)

// TableDao 结构体

type TableDao struct{}

// Create 创建Table记录
// Author
func (dao *TableDao) Create(tx *gorm.DB, table *model.Table) (err error) {
	err = tx.Create(&table).Error
	return err
}

// DeleteByIds 批量删除Table记录
// Author
func (dao *TableDao) DeleteByIds(tx *gorm.DB, ids []string) (err error) {
	err = tx.Delete(&[]model.Table{}, "id in ?", ids).Error
	return err
}

// Update 更新Table记录
// Author
func (dao *TableDao) Update(tx *gorm.DB, table *model.Table) (err error) {
	err = tx.Updates(table).Error
	return err
}

// Get 根据id获取Table记录
// Author
func (dao *TableDao) Get(id string) (err error, table *model.Table) {
	err = global.GormDao.Where("id = ?", id).First(&table).Error
	return
}

// Page 分页获取Table记录
// Author
func (dao *TableDao) Page(param *vo.TablePageView) (err error, page *common.PageInfo) {
	db := global.GormDao.Model(&model.Table{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.TableName != "" {
		db.Where("lower(table_name) like lower(?)", "%"+param.TableName+"%")
	}
	if param.TableComment != "" {
		db.Where("table_comment like ?", "%"+param.TableComment+"%")
	}

	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.Table
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

// List 获取Table记录
// Author
func (dao *TableDao) List(v *vo.TableQueryView) (err error, dataList []*model.Table) {
	db := global.GormDao.Model(&model.Table{})
	// TODO 输入查询条件
	//if data.Id != "" {
	//    db.Where("id = ?", data.Id)
	//}
	// db.Order("create_time desc")
	err = db.Find(&dataList).Error
	return err, dataList
}

// SelectDbTablePage 获取数据库表列表
func (dao *TableDao) SelectDbTablePage(v *vo.TablePageView) (error, *common.PageInfo) {
	tx := global.GormDao.Table("information_schema.tables")
	tx.Select("TABLE_NAME table_name,TABLE_COMMENT table_comment,CREATE_TIME create_time,UPDATE_TIME update_time")
	tx.Where("table_schema = (select database())")
	tx.Where("table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'")
	tx.Where("table_name NOT IN (select table_name from gen_table)")
	if v.TableName != "" {
		tx.Where("lower(table_name) like lower(?)", "%"+v.TableName+"%")
	}
	if v.TableComment != "" {
		tx.Where("table_comment like ?", "%"+v.TableComment+"%")
	}
	tx.Order("create_time desc")
	page := common.CreatePageInfo(v.PageNum, v.PageSize)
	if err := tx.Count(&page.Total).Error; err != nil {
		return err, nil
	}
	tx = tx.Limit(page.Limit).Offset(page.Offset)
	var tmp []*model.Table
	if err := tx.Find(&tmp).Error; err != nil {
		return err, nil
	}
	page.Rows = tmp
	return nil, page
}

func (dao *TableDao) SelectDbTableList(names []string) (error, []*model.Table) {
	var tmp []*model.Table
	tx := global.GormDao.Table("information_schema.tables")
	tx.Select("TABLE_NAME table_name,TABLE_COMMENT table_comment,CREATE_TIME create_time,UPDATE_TIME update_time")
	tx.Where("table_schema = (select database())")
	tx.Where("table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'")
	tx.Where("table_name in (?)", names)
	if err := tx.Find(&tmp).Error; err != nil {
		return err, nil
	}
	return nil, tmp
}

// SelectDbTableByName 根据表名获取数据库表信息
func (dao *TableDao) SelectDbTableByName(name string) (error, *model.Table) {
	var tmp model.Table
	tx := global.GormDao.Model(&model.Table{})
	tx.Where("table_name = ?", name)
	if err := tx.Find(&tmp).Error; err != nil {
		return err, nil
	}
	return nil, &tmp
}
