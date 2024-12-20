// Package dao 自动生成模板 UserDao
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package dao

import (
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/app/admin/sys/service/user/view"
	"matuto-base/src/common"
	"matuto-base/src/global"
)

// UserDao 结构体
type UserDao struct{}

// Create 创建User记录
// Author
func (dao *UserDao) Create(tx *gorm.DB, sysUser model.User) (err error) {
	err = tx.Create(&sysUser).Error
	return err
}

// Delete 删除User记录
// Author
func (dao *UserDao) Delete(tx *gorm.DB, id string) (err error) {
	err = tx.Delete(&[]model.User{}, "id = ?", id).Error
	return err
}

// DeleteByIds 批量删除User记录
// Author
func (dao *UserDao) DeleteByIds(tx *gorm.DB, ids []string) (err error) {
	err = tx.Delete(&[]model.User{}, "id in ?", ids).Error
	return err
}

// Update 更新User记录
// Author
func (dao *UserDao) Update(sysUser model.User) (err error) {
	err = global.GormDao.Updates(&sysUser).Error
	return err
}

// Get 根据id获取User记录
// Author
func (dao *UserDao) Get(id string) (err error, sysUser *model.User) {
	err = global.GormDao.Where("id = ?", id).First(&sysUser).Error
	return
}

// Page 分页获取User记录
// Author
func (dao *UserDao) Page(param *view.UserPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GormDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phone_number, u.sex, u.status, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.Id != "" {
		db.Where("ID = ?", param.Id)
	}
	if param.UserName != "" {
		db.Where("u.user_name like ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		db.Where("u.phone_number like ?", "%"+param.PhoneNumber+"%")
	}
	if param.Status != "" {
		db.Where("u.status = ?", param.Status)
	}
	if param.DeptId != "" {
		db.Where("(u.dept_id = ? or u.dept_id in (SELECT t.id FROM sys_dept t WHERE find_in_set(?, ancestors)))", param.DeptId, param.DeptId)
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var tmp []*model.User
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	page.Rows = tmp
	return err, page
}

// List 获取User记录
// Author
func (dao *UserDao) List(data *model.User) (err error, datas []*model.User) {
	var rows []*model.User
	model := global.GormDao.Table("sys_user u")
	model.Select("distinct u.id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phone_number, u.sex, u.status, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark")
	model.Joins("left join sys_dept d on u.dept_id = d.id")
	if data.Id != "" {
		model = model.Where("ID = ?", data.Id)
	}
	if data.UserName != "" {
		model = model.Where("user_name like ?", "%"+data.UserName+"%")
	}
	if data.PhoneNumber != "" {
		model = model.Where("phone_number like ?", "%"+data.PhoneNumber+"%")
	}
	if data.Status != "" {
		model = model.Where("status = ?", data.Status)
	}
	if data.DeptId != "" {
		model = model.Where("dept_id = ?", data.DeptId)
	}
	if data.DataScopeSql != "" {
		model = model.Where(" ?", data.DataScopeSql)
	}
	model.Order("create_time desc")
	err = model.Find(&rows).Error
	datas = rows
	return err, datas
}

// GetByUserName 根据用户名获取User记录
func (dao *UserDao) GetByUserName(name string) (err error, sysUser *model.User) {
	err = global.GormDao.Where("user_name = ?", name).First(&sysUser).Error
	return
}

// GetByDeptId 根据部门id获取User记录
func (dao *UserDao) GetByDeptId(deptId string) (err error, sysUser []*model.User) {
	err = global.GormDao.Where("dept_id = ?", deptId).Find(&sysUser).Error
	return
}

// CheckFieldUnique 判断指定字段是否为空
func (dao *UserDao) CheckFieldUnique(fieldName, value string) (error, int64) {
	var count int64
	err := global.GormDao.Model(&model.User{}).Where(fieldName+" = ?", value).Count(&count).Error
	return err, count
}

// SelectByField 根据指定字段查询数据
func (dao *UserDao) SelectByField(fieldName string, value string) (error, *model.User) {
	var users []*model.User
	if err := global.GormDao.Model(&model.User{}).Where(fieldName+" = ?", value).Find(&users).Error; err != nil {
		return err, nil
	} else {
		if users != nil && len(users) > 0 {
			return nil, users[0]
		}
		return nil, nil
	}

}

// SelectAllocatedList 获取已分配用户角色的用户列表
func (dao *UserDao) SelectAllocatedList(param *view.UserPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GormDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.user_name, u.nick_name, u.email, u.phone_number, u.status, u.create_time")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	db.Joins("left join sys_user_role ur on u.id = ur.user_id")
	db.Joins("left join sys_role r on ur.role_id = r.id")
	db.Where("r.id = ?", param.RoleId)
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.UserName != "" {
		db.Where("u.user_name like ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		db.Where("u.phone_number like ?", "%"+param.PhoneNumber+"%")
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var tmp []*model.User
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	page.Rows = tmp
	return err, page
}

// SelectUnallocatedList 获取未分配用户角色的用户列表
func (dao *UserDao) SelectUnallocatedList(param *view.UserPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GormDao.Table("sys_user u")
	db.Select("distinct u.id, u.dept_id, u.user_name, u.nick_name, u.email, u.phone_number, u.status, u.create_time")
	db.Joins("left join sys_dept d on u.dept_id = d.id")
	db.Joins("left join sys_user_role ur on u.id = ur.user_id")
	db.Joins("left join sys_role r on ur.role_id = r.id")
	db.Where("(r.id != ? or r.id is null)", param.RoleId)
	db.Where("u.id not in (select u.id from sys_user u inner join sys_user_role ur on u.id = ur.user_id and ur.role_id = ?)", param.RoleId)
	// 如果有条件搜索 下方会自动创建搜索语句
	if param.UserName != "" {
		db.Where("u.user_name like ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		db.Where("u.phone_number like ?", "%"+param.PhoneNumber+"%")
	}
	if param.DataScopeSql != "" {
		db.Where(param.DataScopeSql)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var tmp []*model.User
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&tmp).Error
	page.Rows = tmp
	return err, page
}
