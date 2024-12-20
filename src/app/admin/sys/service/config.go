// Package config 自动生成模板 ConfigService
// @description <TODO description class purpose>
// @author
// @File: config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package service

import (
	"errors"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/common/constants"
	"matuto-base/src/utils/convert"
)

type ConfigService struct {
	configDao dao.ConfigDao
}

// Create 创建Config记录
// Author
func (s *ConfigService) Create(configView *vo.ConfigView) error {
	if err, config := convert.View2Data[vo.ConfigView, model.Config](configView); err != nil {
		return err
	} else {
		return s.configDao.Create(*config)
	}
}

// DeleteByIds 批量删除Config记录
// Author
func (s *ConfigService) DeleteByIds(ids []string) (err error) {
	// 判断是否为系统配置
	for _, id := range ids {
		if err1, config := s.configDao.Get(id); err1 != nil {
			return err1
		} else {
			if config.ConfigType == constants.YES {
				return errors.New("系统内置，不可删除")
			}
		}
	}
	err = s.configDao.DeleteByIds(ids)
	return err
}

// Update 更新Config记录
// Author
func (s *ConfigService) Update(id string, configView *vo.ConfigView) (err error) {
	configView.Id = id
	if err1, config := convert.View2Data[vo.ConfigView, model.Config](configView); err1 != nil {
		return err1
	} else {
		return s.configDao.Update(*config)
	}
}

// Get 根据id获取Config记录
// Author
func (s *ConfigService) Get(id string) (err error, configView *vo.ConfigView) {
	if id == "" {
		return nil, nil
	}
	err1, config := s.configDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, configView = convert.Data2View[vo.ConfigView, model.Config](config)
	return
}

// Page 分页获取Config记录
// Author
func (s *ConfigService) Page(pageInfo *vo.ConfigPageView) (err error, res *common.PageInfo) {
	err, res = s.configDao.Page(pageInfo)
	if err != nil {
		return err, nil
	}
	return convert.PageData2ViewList[vo.ConfigView, model.Config](res)
}

func (s *ConfigService) List(v *vo.ConfigView) (err error, views []*vo.ConfigView) {
	err, data := convert.View2Data[vo.ConfigView, model.Config](v)
	if err != nil {
		return err, nil
	}
	var datas []*model.Config
	if err, datas = s.configDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = convert.Data2ViewList[vo.ConfigView, model.Config](datas)
		return
	}
}

// SelectConfigByKey 根据key查询Config记录
func (s *ConfigService) SelectConfigByKey(key string) (error, *vo.ConfigView) {
	if err, config := s.configDao.SelectConfigByKey(key); err != nil {
		return err, nil
	} else {
		if err, configView := convert.Data2View[vo.ConfigView, model.Config](config); err != nil {
			return err, nil
		} else {
			return nil, configView
		}
	}
}
