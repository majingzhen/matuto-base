// Package dao 自动生成模板 TableColumnDao
// @description <TODO description class purpose>
// @author
// @File: gen_table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/gen/model"
	"matuto-base/src/global"
)

// TableColumnDao 结构体

type TableColumnDao struct{}

// Create 创建TableColumn记录
// Author
func (dao *TableColumnDao) Create(tx *gorm.DB, genTableColumn *model.TableColumn) (err error) {
	err = tx.Create(genTableColumn).Error
	return err
}

// DeleteByIds 批量删除TableColumn记录
// Author
func (dao *TableColumnDao) DeleteByIds(tx *gorm.DB, ids []string) (err error) {
	err = tx.Delete(&[]model.TableColumn{}, "id in ?", ids).Error
	return err
}

// Update 更新TableColumn记录
// Author
func (dao *TableColumnDao) Update(tx *gorm.DB, genTableColumn *model.TableColumn) (err error) {
	err = tx.Updates(genTableColumn).Error
	return err
}

// Get 根据id获取TableColumn记录
// Author
func (dao *TableColumnDao) Get(id string) (err error, genTableColumn *model.TableColumn) {
	err = global.GormDao.Where("id = ?", id).First(&genTableColumn).Error
	return
}

func (dao *TableColumnDao) SelectDbTableColumns(tx *gorm.DB, tableName string) (err error, dataList []*model.TableColumn) {
	db := tx.Table("information_schema.columns")
	db.Select("COLUMN_NAME column_name,( CASE WHEN ( IS_NULLABLE = 'no' && COLUMN_KEY != 'PRI' ) THEN '1' ELSE NULL END ) AS is_required,( CASE WHEN COLUMN_KEY = 'PRI' THEN '1' ELSE '0' END ) AS is_pk,ORDINAL_POSITION AS sort,COLUMN_COMMENT column_comment,( CASE WHEN extra = 'AUTO_INCREMENT' THEN '1' ELSE '0' END ) AS is_increment,DATA_TYPE data_type, COLUMN_TYPE column_type, CHARACTER_MAXIMUM_LENGTH column_length")
	db.Where("table_schema = (SELECT DATABASE())")
	db.Where("table_name = ?", tableName).Find(&dataList)
	err = db.Find(&dataList).Error
	return err, dataList
}

// DeleteByTableIds 根据表id删除列信息
func (dao *TableColumnDao) DeleteByTableIds(tx *gorm.DB, tableIds []string) error {
	return tx.Table("gen_table_column").Where("table_id in (?)", tableIds).Delete(&model.TableColumn{}).Error
}

// GetColumnListByTableId 根据表id获取列信息
func (dao *TableColumnDao) GetColumnListByTableId(tableId string) (error, []*model.TableColumn) {
	var columns []*model.TableColumn
	err := global.GormDao.Table("gen_table_column").Where("table_id = ?", tableId).Order("sort").Find(&columns).Error
	return err, columns
}

// SelectPkColumn 根据表id获取主键列
func (dao *TableColumnDao) SelectPkColumn(tableId string) (err error, genTableColumn *model.TableColumn) {
	tx := global.GormDao.Table("gen_table_column")
	tx.Where("table_id = ?", tableId)
	tx.Where("is_pk = ?", "1")
	err = tx.First(&genTableColumn).Error
	return
}

// SelectSearchColumn 根据表id获取搜索列
func (dao *TableColumnDao) SelectSearchColumn(id string) (error, []*model.TableColumn) {
	var columns []*model.TableColumn
	err := global.GormDao.Table("gen_table_column").Where("table_id = ?", id).Where("is_query = ?", "1").Order("sort").Find(&columns).Error
	return err, columns

}

// SelectInsertColumn 根据表id获取插入列
func (dao *TableColumnDao) SelectInsertColumn(id string) (error, []*model.TableColumn) {
	var columns []*model.TableColumn
	err := global.GormDao.Table("gen_table_column").Where("table_id = ?", id).Where("is_insert = ?", "1").Order("sort").Find(&columns).Error
	return err, columns
}

// SelectEditColumn 根据表id获取编辑列
func (dao *TableColumnDao) SelectEditColumn(id string) (error, []*model.TableColumn) {
	var columns []*model.TableColumn
	err := global.GormDao.Table("gen_table_column").Where("table_id = ?", id).Where("is_edit = ?", "1").Order("sort").Find(&columns).Error
	return err, columns
}

// SelectListColumn 根据表id获取列表列
func (dao *TableColumnDao) SelectListColumn(id string) (error, []*model.TableColumn) {
	var columns []*model.TableColumn
	err := global.GormDao.Table("gen_table_column").Where("table_id = ?", id).Where("is_list = ?", "1").Order("sort").Find(&columns).Error
	return err, columns
}
