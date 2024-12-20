// Package sys_dict_data 自动生成模板 DictDataService
// @description <TODO description class purpose>
// @author
// @File: sys_dict_data
// @version 1.0.0
// @create 2023-08-20 19:08:42
package dict_data

import (
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/service/dict_data/view"
	"matuto-base/src/common"
)

type DictDataService struct {
	sysDictDataDao dao.DictDataDao
	viewUtils      view.DictDataViewUtils
}

// Create 创建DictData记录
// Author
func (s *DictDataService) Create(sysDictDataView *view.DictDataView) (err error) {
	err1, sysDictData := s.viewUtils.View2Data(sysDictDataView)
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
func (s *DictDataService) Update(id string, sysDictDataView *view.DictDataView) (err error) {
	sysDictDataView.Id = id
	err1, sysDictData := s.viewUtils.View2Data(sysDictDataView)
	if err1 != nil {
		return err1
	}
	err = s.sysDictDataDao.Update(*sysDictData)
	return err
}

// Get 根据id获取DictData记录
// Author
func (s *DictDataService) Get(id string) (err error, sysDictDataView *view.DictDataView) {
	if id == "" {
		return nil, nil
	}
	err1, sysDictData := s.sysDictDataDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err2, sysDictDataView := s.viewUtils.Data2View(sysDictData)
	if err2 != nil {
		return err2, nil
	}
	return
}

// Page 分页获取DictData记录
// Author
func (s *DictDataService) Page(pageInfo *view.DictDataPageView) (err error, res *common.PageInfo) {
	if err, res = s.sysDictDataDao.Page(pageInfo); err != nil {
		return err, nil
	}
	return s.viewUtils.PageData2ViewList(res)
}

// GetByType 根据类型获取数据
func (s *DictDataService) GetByType(dictType string) (err error, views []*view.DictDataView) {
	err1, datas := s.sysDictDataDao.GetByType(dictType)
	if err1 != nil {
		return err1, nil
	}
	err2, views := s.viewUtils.Data2ViewList(datas)
	if err2 != nil {
		return err2, nil
	}
	return
}
