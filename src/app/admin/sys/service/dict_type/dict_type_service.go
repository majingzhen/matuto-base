// Package dict_type 自动生成模板 DictTypeService
// @description <TODO description class purpose>
// @author
// @File: dict_type
// @version 1.0.0
// @create 2023-08-18 13:41:26
package dict_type

import (
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/utils/convert"
)

type DictTypeService struct {
	sysDictTypeDao dao.DictTypeDao
}

// Create 创建DictType记录
// Author
func (s *DictTypeService) Create(sysDictTypeView *vo.DictTypeView) (err error) {
	err1, sysDictType := convert.View2Data[vo.DictTypeView, model.DictType](sysDictTypeView)
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
func (s *DictTypeService) Update(id string, sysDictTypeView *vo.DictTypeView) (err error) {
	sysDictTypeView.Id = id
	err1, sysDictType := convert.View2Data[vo.DictTypeView, model.DictType](sysDictTypeView)
	if err1 != nil {
		return err1
	}
	err = s.sysDictTypeDao.Update(*sysDictType)
	return err
}

// Get 根据id获取DictType记录
// Author
func (s *DictTypeService) Get(id string) (err error, sysDictTypeView *vo.DictTypeView) {
	if id == "" {
		return nil, nil
	}
	err1, sysDictType := s.sysDictTypeDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictTypeView := convert.Data2View[vo.DictTypeView, model.DictType](sysDictType)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取DictType记录
// Author
func (s *DictTypeService) Page(pageInfo *vo.DictTypePageView) (err error, res *common.PageInfo) {
	if err, res = s.sysDictTypeDao.Page(pageInfo); err != nil {
		return err, nil
	}
	return convert.PageData2ViewList[vo.DictTypeView, model.DictType](res)
}

// SelectDictTypeAll 获取全部数据
func (s *DictTypeService) SelectDictTypeAll() (err error, views []*vo.DictTypeView) {
	err, datas := s.sysDictTypeDao.SelectDictTypeAll()
	if err != nil {
		return err, nil
	}
	err, views = convert.Data2ViewList[vo.DictTypeView, model.DictType](datas)
	if err != nil {
		return err, nil
	}
	return
}
