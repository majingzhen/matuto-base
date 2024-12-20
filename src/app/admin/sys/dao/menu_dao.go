// Package dao 自动生成模板 MenuDao
// @description <TODO description class purpose>
// @author
// @File: menu
// @version 1.0.0
// @create 2023-08-18 13:41:26
package dao

import (
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common/constants"
	"matuto-base/src/global"
)

// MenuDao 结构体

type MenuDao struct{}

// Create 创建Menu记录
// Author
func (dao *MenuDao) Create(sysMenu model.Menu) (err error) {
	err = global.GormDao.Create(&sysMenu).Error
	return err
}

// Delete 删除Menu记录
// Author
func (dao *MenuDao) Delete(id string) (err error) {
	err = global.GormDao.Delete(&model.Menu{}, "id = ?", id).Error
	return err
}

// Update 更新Menu记录
// Author
func (dao *MenuDao) Update(sysMenu model.Menu) (err error) {
	err = global.GormDao.Updates(&sysMenu).Error
	return err
}

// Get 根据id获取Menu记录
// Author
func (dao *MenuDao) Get(id string) (err error, sysMenu *model.Menu) {
	err = global.GormDao.Where("id = ?", id).First(&sysMenu).Error
	return
}

// GetMenuPermissionByRoleId 根据角色id查询菜单权限
func (dao *MenuDao) GetMenuPermissionByRoleId(roleId string) (err error, perms []string) {
	var rows []model.Menu
	db := global.GormDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_role r on r.id = rm.role_id")
	db.Select("distinct m.perms")
	db.Where("rm.role_id = ? and r.status = ?", roleId, constants.STATUS_NORMAL)
	err = db.Scan(&rows).Error
	if err != nil {
		return
	}
	for _, menu := range rows {
		perms = append(perms, menu.Perms)
	}
	return err, perms
}

// GetMenuPermissionByUserId 根据用户id查询菜单权限
func (dao *MenuDao) GetMenuPermissionByUserId(userId string) (err error, perms []string) {
	var rows []model.Menu
	db := global.GormDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_user_role ur on rm.role_id = ur.role_id")
	db.Select("distinct m.perms")
	db.Where("ur.user_id = ? and r.status = ? and m.status = ?", userId, constants.STATUS_NORMAL, constants.STATUS_NORMAL)
	err = db.Scan(&rows).Error
	if err != nil {
		return
	}
	for _, menu := range rows {
		perms = append(perms, menu.Perms)
	}
	return err, perms
}

// SelectMenuAll 查询所有菜单
func (dao *MenuDao) SelectMenuAll() (err error, menus []*model.Menu) {
	db := global.GormDao.Model(&[]model.Menu{})
	db.Where("status = ? and menu_type in (?, ?)", constants.STATUS_NORMAL, constants.MENU_TYPE_DIR, constants.MENU_TYPE_MENU)
	err1 := db.Find(&menus).Error
	if err1 != nil {
		return err1, nil
	}
	return err, menus
}

// SelectMenuByUserId 根据用户id查询菜单
func (dao *MenuDao) SelectMenuByUserId(userId string) (err error, menus []*model.Menu) {
	db := global.GormDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_user_role ur on rm.role_id = ur.role_id")
	db.Joins("left join sys_role r on r.id = rm.role_id")
	db.Select("distinct m.id, m.parent_id, m.menu_name, m.path, m.component, m.`query`, m.visible, m.status, perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time")
	db.Where("ur.user_id = ? and r.status = ? and m.status = ? and menu_type in (?, ?)", userId, constants.STATUS_NORMAL, constants.STATUS_NORMAL, constants.MENU_TYPE_DIR, constants.MENU_TYPE_MENU)
	db.Order("m.parent_id, m.order_num")
	err = db.Scan(&menus).Error
	if err != nil {
		return err, nil
	}
	return err, menus
}

// SelectMenuList 查询Menu记录
func (dao *MenuDao) SelectMenuList(data *model.Menu) (err error, menus []*model.Menu) {
	db := global.GormDao.Table(data.TableName())
	db.Select("distinct id, parent_id, menu_name, path, component, `query`, visible, status, perms, is_frame, is_cache, menu_type, icon, order_num, create_time")
	if data.MenuName != "" {
		db.Where("menu_name like ?", "%"+data.MenuName+"%")
	}
	if data.Visible != "" {
		db.Where("visible = ?", data.Visible)
	}
	if data.Status != "" {
		db.Where("status = ?", data.Status)
	}
	db.Order("parent_id, order_num")
	err = db.Scan(&menus).Error
	return err, menus
}

// SelectMenuListByUserId 根据用户id查询菜单
func (dao *MenuDao) SelectMenuListByUserId(data *model.Menu, userId string) (err error, menus []*model.Menu) {
	db := global.GormDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Joins("left join sys_user_role ur on rm.role_id = ur.role_id")
	db.Select("distinct m.id, m.parent_id, m.menu_name, m.path, m.component, m.`query`, m.visible, m.status, m.perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time")
	db.Where("ur.user_id = ? ", userId)
	if data.MenuName != "" {
		db.Where("menu_name like ?", "%"+data.MenuName+"%")
	}
	if data.Visible != "" {
		db.Where("visible = ?", data.Visible)
	}
	if data.Status != "" {
		db.Where("status = ?", data.Status)
	}
	db.Order("m.parent_id, m.order_num")
	err = db.Scan(&menus).Error
	return err, menus
}

// CheckMenuNameUniqueAll 判断名称是否唯一
func (dao *MenuDao) CheckMenuNameUniqueAll(data *model.Menu) (error, bool) {
	var menu model.Menu
	db := global.GormDao.Model(&model.Menu{})
	db.Where("menu_name = ?", data.MenuName)
	if data.ParentId != "" {
		db.Where("parent_id = ?", data.ParentId)
	}
	err := db.First(&menu).Error
	if err != nil {
		return err, false
	}
	return nil, true
}

// SelectMenuByParentId 根据父级id查询菜单
func (dao *MenuDao) SelectMenuByParentId(id string) (err error, rows []*model.Menu) {
	db := global.GormDao.Model(&model.Menu{})
	db.Where("parent_id = ?", id)
	err = db.Find(&rows).Error
	if err != nil {
		return err, nil
	}
	return nil, rows
}

// CheckMenuExistRole 判断菜单是否存在角色
func (dao *MenuDao) CheckMenuExistRole(roleId string) (error, bool) {
	var count int64
	db := global.GormDao.Table("sys_role_menu").Where("menu_id = ?", roleId)
	err := db.Count(&count).Error
	if err != nil {
		return err, false
	}
	if count > 0 {
		return nil, true
	}
	return nil, false
}

// SelectMenuListByRoleId 根据角色id查询菜单
func (dao *MenuDao) SelectMenuListByRoleId(id string, strictly bool) (error, []string) {
	var rows []model.Menu
	db := global.GormDao.Table("sys_menu m")
	db.Joins("left join sys_role_menu rm on m.id = rm.menu_id")
	db.Select("m.id")
	db.Where("rm.role_id = ?", id)
	if strictly {
		db.Where("m.id not in (select m.parent_id from sys_menu m inner join sys_role_menu rm on m.id = rm.menu_id and rm.role_id = ?)", id)
	}
	db.Order(" m.parent_id, m.order_num")
	err := db.Scan(&rows).Error
	if err != nil {
		return err, nil
	}
	var perms []string
	for _, menu := range rows {
		perms = append(perms, menu.Id)
	}
	return nil, perms
}
