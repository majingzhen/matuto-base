// Package sys_dict_data 自动生成模板 DictDataService
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package service

import (
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/utils/convert"
)

type DictDataService struct {
	sysDictDataDao dao.DictDataDao
}

// Create 创建DictData记录
// Author
func (s *DictDataService) Create(sysDictDataView *vo.DictDataView) (err error) {
	err1, sysDictData := convert.View2Data[vo.DictDataView, model.DictData](sysDictDataView)
	if err1 != nil {
		return err1
	}
	err2 := s.sysDictDataDao.Create(*sysDictData)
	if err2 != nil {
		return err2
	}
	return nil
}

// DeleteByIds 批量删除DictData记录
// Author
func (s *DictDataService) DeleteByIds(ids []string) (err error) {
	err = s.sysDictDataDao.DeleteByIds(ids)
	return err
}

// Update 更新DictData记录
// Author
func (s *DictDataService) Update(id string, sysDictDataView *vo.DictDataView) (err error) {
	sysDictDataView.Id = id
	err1, sysDictData := convert.View2Data[vo.DictDataView, model.DictData](sysDictDataView)
	if err1 != nil {
		return err1
	}
	err = s.sysDictDataDao.Update(*sysDictData)
	return err
}

// Get 根据id获取DictData记录
// Author
func (s *DictDataService) Get(id string) (err error, sysDictDataView *vo.DictDataView) {
	if id == "" {
		return nil, nil
	}
	err1, sysDictData := s.sysDictDataDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictDataView := convert.Data2View[vo.DictDataView, model.DictData](sysDictData)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取DictData记录
// Author
func (s *DictDataService) Page(pageInfo *vo.DictDataPageView) (err error, res *common.PageInfo) {
	if err, res = s.sysDictDataDao.Page(pageInfo); err != nil {
		return err, nil
	}
	return convert.PageData2ViewList[vo.DictDataView, model.DictData](res)
}

// GetByType 根据类型获取数据
func (s *DictDataService) GetByType(dictType string) (err error, views []*vo.DictDataView) {
	err1, datas := s.sysDictDataDao.GetByType(dictType)
	if err1 != nil {
		return err1, nil
	}
	err2, views := convert.Data2ViewList[vo.DictDataView, model.DictData](datas)
	if err2 != nil {
		return err2, nil
	}
	return
}
