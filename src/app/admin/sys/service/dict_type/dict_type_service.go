// Package dict_type 自动生成模板 DictTypeService
// @description <TODO description class purpose>
// @author
// @File: dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package dict_type

import (
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/service/dict_type/view"
	"matuto-base/src/common"
)

type DictTypeService struct {
	sysDictTypeDao dao.DictTypeDao
	viewUtils      view.DictTypeViewUtils
}

// Create 创建DictType记录
// Author
func (s *DictTypeService) Create(sysDictTypeView *view.DictTypeView) (err error) {
	err1, sysDictType := s.viewUtils.View2Data(sysDictTypeView)
	if err1 != nil {
		return err1
	}
	err2 := s.sysDictTypeDao.Create(*sysDictType)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete 删除DictType记录
// Author
func (s *DictTypeService) Delete(id string) (err error) {
	err = s.sysDictTypeDao.Delete(id)
	return err
}

// DeleteByIds 批量删除DictType记录
// Author
func (s *DictTypeService) DeleteByIds(ids []string) (err error) {
	err = s.sysDictTypeDao.DeleteByIds(ids)
	return err
}

// Update 更新DictType记录
// Author
func (s *DictTypeService) Update(id string, sysDictTypeView *view.DictTypeView) (err error) {
	sysDictTypeView.Id = id
	err1, sysDictType := s.viewUtils.View2Data(sysDictTypeView)
	if err1 != nil {
		return err1
	}
	err = s.sysDictTypeDao.Update(*sysDictType)
	return err
}

// Get 根据id获取DictType记录
// Author
func (s *DictTypeService) Get(id string) (err error, sysDictTypeView *view.DictTypeView) {
	if id == "" {
		return nil, nil
	}
	err1, sysDictType := s.sysDictTypeDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictTypeView := s.viewUtils.Data2View(sysDictType)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取DictType记录
// Author
func (s *DictTypeService) Page(pageInfo *view.DictTypePageView) (err error, res *common.PageInfo) {
	if err, res = s.sysDictTypeDao.Page(pageInfo); err != nil {
		return err, nil
	}
	return s.viewUtils.PageData2ViewList(res)
}

// SelectDictTypeAll 获取全部数据
func (s *DictTypeService) SelectDictTypeAll() (err error, views []*view.DictTypeView) {
	err, datas := s.sysDictTypeDao.SelectDictTypeAll()
	if err != nil {
		return err, nil
	}
	err, views = s.viewUtils.Data2ViewList(datas)
	if err != nil {
		return err, nil
	}
	return
}
