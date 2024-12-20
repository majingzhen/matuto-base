// Package table_column 自动生成模板 TableColumnService
// @description <TODO description class purpose>
// @author
// @File: gen_table_column
// @version 1.0.0
// @create 2023-08-31 09:09:53
package table_column

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/gen/dao"
	"matuto-base/src/app/admin/gen/service/table_column/view"
	"matuto-base/src/global"
)

type TableColumnService struct {
	tableColumnDao dao.TableColumnDao
	viewUtils      view.TableColumnViewUtils
}

// DeleteByIds 批量删除TableColumn记录
// Author
func (s *TableColumnService) DeleteByIds(ids []string) (err error) {
	err = s.tableColumnDao.DeleteByIds(global.GormDao, ids)
	return err
}

// Update 更新TableColumn记录
// Author
func (s *TableColumnService) Update(tableColumnView *view.TableColumnView, tx ...*gorm.DB) error {
	if err, tableColumn := s.viewUtils.View2Data(tableColumnView); err != nil {
		return err
	} else {
		if tx != nil {
			return s.tableColumnDao.Update(tx[0], tableColumn)
		} else {
			return s.tableColumnDao.Update(global.GormDao, tableColumn)
		}
	}
}

// Get 根据id获取TableColumn记录
// Author
func (s *TableColumnService) Get(id string) (err error, tableColumnView *view.TableColumnView) {
	err1, tableColumn := s.tableColumnDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, tableColumnView = s.viewUtils.Data2View(tableColumn)
	return
}

// GetColumnListByTableId 根据tableId获取TableColumn列表
func (s *TableColumnService) GetColumnListByTableId(tableId string) (error, []*view.TableColumnView) {
	if err, dataList := s.tableColumnDao.GetColumnListByTableId(tableId); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}

func (s *TableColumnService) SelectPkColumn(tableId string) (error, *view.TableColumnView) {
	if err, column := s.tableColumnDao.SelectPkColumn(tableId); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2View(column)
	}
}

// SelectSearchColumn 根据tableId获取TableColumn列表
func (s *TableColumnService) SelectSearchColumn(tableId string) (error, []*view.TableColumnView) {
	if err, dataList := s.tableColumnDao.SelectSearchColumn(tableId); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}

// SelectInsertColumn 根据tableId获取TableColumn列表
func (s *TableColumnService) SelectInsertColumn(id string) (error, []*view.TableColumnView) {
	if err, dataList := s.tableColumnDao.SelectInsertColumn(id); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}

// SelectEditColumn 根据tableId获取TableColumn列表
func (s *TableColumnService) SelectEditColumn(id string) (error, []*view.TableColumnView) {
	if err, dataList := s.tableColumnDao.SelectEditColumn(id); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}

}

// SelectListColumn 根据tableId获取TableColumn列表
func (s *TableColumnService) SelectListColumn(id string) (error, []*view.TableColumnView) {
	if err, dataList := s.tableColumnDao.SelectListColumn(id); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}
