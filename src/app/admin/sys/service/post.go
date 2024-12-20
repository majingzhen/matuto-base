// Package service 自动生成模板 PostService
// @description <TODO description class purpose>
// @author
// @File: sys_post
// @version 1.0.0
// @create 2023-08-21 17:37:56
package service

import (
	"errors"
	"fmt"
	view2 "matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/utils/convert"
)

type PostService struct {
	sysPostDao dao.PostDao
}

// Create 创建Post记录
// Author
func (s *PostService) Create(sysPostView *view2.PostView) error {
	// 校验是否重复
	if err := s.CheckPostCodeUnique(sysPostView.PostCode); err != nil {
		return err
	}
	if err := s.CheckPostNameUnique(sysPostView.PostName); err != nil {
		return err
	}
	if err, sysPost := convert.View2Data[view2.PostView, model.Post](sysPostView); err != nil {
		return err
	} else {
		return s.sysPostDao.Create(*sysPost)
	}
}

// DeleteByIds 批量删除Post记录
// Author
func (s *PostService) DeleteByIds(ids []string) error {
	for _, id := range ids {
		err, postView := s.Get(id)
		if err != nil {
			return err
		}
		if err1, count := s.sysPostDao.CheckPostExistUser(id); err1 != nil {
			return err1
		} else {
			if count > 0 {
				return errors.New(fmt.Sprintf("%s 已分配,不能删除", postView.PostName))
			}
		}
	}
	return s.sysPostDao.DeleteByIds(ids)
}

// Update 更新Post记录
// Author
func (s *PostService) Update(id string, sysPostView *view2.PostView) error {
	// 校验是否重复
	if err := s.CheckPostCodeUnique(sysPostView.PostCode); err != nil {
		return err
	}
	if err := s.CheckPostNameUnique(sysPostView.PostName); err != nil {
		return err
	}
	sysPostView.Id = id
	if err, sysPost := convert.View2Data[view2.PostView, model.Post](sysPostView); err != nil {
		return err
	} else {
		return s.sysPostDao.Update(*sysPost)
	}
}

// Get 根据id获取Post记录
// Author
func (s *PostService) Get(id string) (err error, sysPostView *view2.PostView) {
	if id == "" {
		return nil, nil
	}
	err1, sysPost := s.sysPostDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, sysPostView = convert.Data2View[view2.PostView, model.Post](sysPost)
	return
}

// Page 分页获取Post记录
// Author
func (s *PostService) Page(pageInfo *view2.PostPageView) (err error, res *common.PageInfo) {
	if err, res = s.sysPostDao.Page(pageInfo); err != nil {
		return err, nil
	}
	return convert.PageData2ViewList[view2.PostView, model.Post](res)
}

// List 获取Post列表
// Author
func (s *PostService) List(v *view2.PostView) (err error, views []*view2.PostView) {
	err, data := convert.View2Data[view2.PostView, model.Post](v)
	if err != nil {
		return err, nil
	}
	var datas []*model.Post
	if err, datas = s.sysPostDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = convert.Data2ViewList[view2.PostView, model.Post](datas)
		return
	}
}

// CheckPostCodeUnique 校验岗位编码是否唯一
// Author
func (s *PostService) CheckPostCodeUnique(postCode string) error {
	if err, count := s.sysPostDao.CheckPostCodeUnique(postCode); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("岗位编码已存在")
		}
	}
	return nil
}

// CheckPostNameUnique 校验岗位名称是否唯一
// Author
func (s *PostService) CheckPostNameUnique(postName string) error {
	if err, count := s.sysPostDao.CheckPostNameUnique(postName); err != nil {
		return err
	} else {
		if count > 0 {
			return errors.New("岗位名称已存在")
		}
	}
	return nil
}

func (s *PostService) SelectPostAll() (err error, views []*view2.PostView) {
	err, views = s.List(&view2.PostView{})
	return
}

func (s *PostService) SelectPostIdListByUserId(userId string) (err error, ids []string) {
	err, dataList := s.sysPostDao.SelectPostListByUserId(userId)
	if err != nil {
		return err, nil
	}
	for _, data := range dataList {
		ids = append(ids, data.Id)
	}
	return
}

// SelectPostListByUserId 根据用户ID查询岗位
func (s *PostService) SelectPostListByUserId(userId string) (err error, views []*view2.PostView) {
	err, dataList := s.sysPostDao.SelectPostListByUserId(userId)
	if err != nil {
		return err, nil
	}
	err, views = convert.Data2ViewList[view2.PostView, model.Post](dataList)
	return
}
