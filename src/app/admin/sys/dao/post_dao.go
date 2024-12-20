// Package dao 自动生成模板 PostDao
// @description <TODO description class purpose>
// @author
// @File: post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package dao

import (
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/global"
)

// PostDao 结构体

type PostDao struct{}

// Create 创建Post记录
// Author
func (dao *PostDao) Create(sysPost model.Post) (err error) {
	err = global.GormDao.Create(&sysPost).Error
	return err
}

// DeleteByIds 批量删除Post记录
// Author
func (dao *PostDao) DeleteByIds(ids []string) (err error) {
	err = global.GormDao.Delete(&[]model.Post{}, "id in ?", ids).Error
	return err
}

// Update 更新Post记录
// Author
func (dao *PostDao) Update(sysPost model.Post) (err error) {
	err = global.GormDao.Updates(&sysPost).Error
	return err
}

// Get 根据id获取Post记录
// Author
func (dao *PostDao) Get(id string) (err error, sysPost *model.Post) {
	err = global.GormDao.Where("id = ?", id).First(&sysPost).Error
	return
}

// Page 分页获取Post记录
// Author
func (dao *PostDao) Page(param *vo.PostPageView) (err error, page *common.PageInfo) {
	// 创建model
	db := global.GormDao.Model(&model.Post{})
	if param.PostName != "" {
		db.Where("post_name like ?", "%"+param.PostName+"%")
	}
	if param.PostCode != "" {
		db.Where("post_code like ?", "%"+param.PostCode+"%")
	}
	if param.Status != "" {
		db.Where("status = ?", param.Status)
	}
	page = common.CreatePageInfo(param.PageNum, param.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	// 生成排序信息
	if param.OrderByColumn != "" {
		db.Order(param.OrderByColumn + " " + param.IsAsc + " ")
	}
	var dataList []*model.Post
	err = db.Limit(page.Limit).Offset(page.Offset).Find(&dataList).Error
	page.Rows = dataList
	return err, page
}

// List 获取Post记录
// Author
func (dao *PostDao) List(data *model.Post) (err error, datas []*model.Post) {
	db := global.GormDao.Model(&model.Post{})
	if data.PostName != "" {
		db.Where("post_name like ?", "%"+data.PostName+"%")
	}
	if data.PostCode != "" {
		db.Where("post_code like ?", "%"+data.PostCode+"%")
	}
	if data.Status != "" {
		db.Where("status = ?", data.Status)
	}
	db.Order("create_time desc")
	err = db.Find(&datas).Error
	return err, datas
}

// CheckPostNameUnique 校验岗位名称是否唯一
// Author
func (dao *PostDao) CheckPostNameUnique(postName string) (err error, count int64) {
	var sysPost model.Post
	err = global.GormDao.Model(&model.Post{}).Where("post_name = ?", postName).First(&sysPost).Count(&count).Error
	return
}

// CheckPostCodeUnique 校验岗位编码是否唯一
// Author
func (dao *PostDao) CheckPostCodeUnique(postCode string) (err error, count int64) {
	var sysPost model.Post
	err = global.GormDao.Model(&model.Post{}).Where("post_code = ?", postCode).First(&sysPost).Count(&count).Error
	return
}

func (dao *PostDao) CheckPostExistUser(postId string) (err error, count int64) {
	err = global.GormDao.Table("sys_user_post").Where("post_id = ?", postId).Count(&count).Error
	return
}

// SelectPostListByUserId 根据用户ID查询岗位
func (dao *PostDao) SelectPostListByUserId(userId string) (error, []*model.Post) {
	db := global.GormDao.Table("sys_post p")
	db.Joins("left join sys_user_post up on p.id = up.post_id")
	db.Joins("left join sys_user u on u.id = up.user_id")
	db.Where("u.id = ?", userId)
	db.Select("distinct p.id, p.post_code, p.post_name, p.post_sort, p.status, p.create_by, p.create_time, p.remark ")
	var res []*model.Post
	err := db.Find(&res).Error
	return err, res
}
