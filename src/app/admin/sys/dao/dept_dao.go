// Package dao 自动生成模板 DeptDao
// @description <TODO description class purpose>
// @author
// @File: dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/global"
)

// DeptDao 结构体

type DeptDao struct{}

// Create 创建Dept记录
// Author
func (dao *DeptDao) Create(sysDept model.Dept) (err error) {
	err = global.GormDao.Create(&sysDept).Error
	return err
}

// DeleteByIds 批量删除Dept记录
// Author
func (dao *DeptDao) DeleteByIds(ids []string) (err error) {
	err = global.GormDao.Delete(&[]model.Dept{}, "id in ?", ids).Error
	return err
}

// Update 更新Dept记录
// Author
func (dao *DeptDao) Update(tx *gorm.DB, sysDept model.Dept) (err error) {
	err = tx.Updates(&sysDept).Error
	return err
}

// Get 根据id获取Dept记录
// Author
func (dao *DeptDao) Get(id string) (err error, sysDept *model.Dept) {
	err = global.GormDao.Table("sys_dept").Where("id = ?", id).First(&sysDept).Error
	return
}

// List 获取Dept记录
// Author
func (dao *DeptDao) List(data *model.Dept) (err error, datas []*model.Dept) {
	db := global.GormDao.Model(&model.Dept{}).Table("sys_dept d")
	// TODO 输入查询条件
	if data.DeptName != "" {
		db.Where("d.dept_name like ?", "?"+data.DeptName+"?")
	}
	if data.Status != "" {
		db.Where("d.status = ?", data.Status)
	}
	if data.ParentId != "" {
		db.Where("d.parent_id = ?", data.ParentId)
	}
	if data.Id != "" {
		db.Where("d.id = ?", data.Id)
	}
	if data.DataScopeSql != "" {
		db.Where(data.DataScopeSql)
	}
	db.Order("d.parent_id, d.order_num")
	err = db.Find(&datas).Error
	return err, datas
}

// CheckDeptNameALL 检查部门名称是否存在
// Author
func (dao *DeptDao) CheckDeptNameALL(deptName, parentId string) (err error, res bool) {
	var count int64
	err = global.GormDao.Model(&model.Dept{}).Where("dept_name = ?,and parent_id = ?", deptName, parentId).Count(&count).Error
	if err != nil {
		return err, false
	}
	return err, count > 0
}

// SelectChildrenDeptById 根据id查询所有下级部门
func (dao *DeptDao) SelectChildrenDeptById(id string) (err error, res *[]model.Dept) {
	var rows []model.Dept
	err = global.GormDao.Model(&model.Dept{}).Where(" find_in_set(?, ancestors)", id).Find(&rows).Error
	res = &rows
	return
}

// Delete 删除Dept记录
func (dao *DeptDao) Delete(id string) error {
	return global.GormDao.Delete(&model.Dept{}, "id = ?", id).Error
}

// SelectDeptListByRoleId 根据角色id查询部门id
func (dao *DeptDao) SelectDeptListByRoleId(id string, strictly bool) (error, []string) {
	db := global.GormDao.Table("sys_dept d")
	db.Joins("left join sys_role_dept rd on d.id = rd.dept_id")
	db.Where("rd.role_id = ?", id)
	if strictly {
		db.Where("d.id not in (select d.parent_id from sys_dept d inner join sys_role_dept rd on d.id = rd.dept_id and rd.role_id = ?)", strictly)
	}
	db.Order("d.parent_id, d.order_num")
	var rows []model.Dept
	err := db.Find(&rows).Error
	if err != nil {
		return err, nil
	} else {
		var ids []string
		for _, row := range rows {
			ids = append(ids, row.Id)
		}
		return nil, ids
	}
}
