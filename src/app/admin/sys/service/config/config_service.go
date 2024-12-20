// Package config 自动生成模板 ConfigService
// @description <TODO description class purpose>
// @author
// @File: config
// @version 1.0.0
// @create 2023-08-21 14:20:37
package config

import (
	"errors"
	"matuto-base/src/app/admin/sys/dao"
	model2 "matuto-base/src/app/admin/sys/model"
	"matuto-base/src/app/admin/sys/service/config/view"
	"matuto-base/src/common"
	"matuto-base/src/common/constants"
)

type ConfigService struct {
	configDao dao.ConfigDao
	viewUtils view.ConfigViewUtils
}

// Create 创建Config记录
// Author
func (s *ConfigService) Create(configView *view.ConfigView) error {
	if err, config := s.viewUtils.View2Data(configView); err != nil {
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
func (s *ConfigService) Update(id string, configView *view.ConfigView) (err error) {
	configView.Id = id
	if err1, config := s.viewUtils.View2Data(configView); err1 != nil {
		return err1
	} else {
		return s.configDao.Update(*config)
	}
}

// Get 根据id获取Config记录
// Author
func (s *ConfigService) Get(id string) (err error, configView *view.ConfigView) {
	if id == "" {
		return nil, nil
	}
	err1, config := s.configDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, configView = s.viewUtils.Data2View(config)
	return
}

// Page 分页获取Config记录
// Author
func (s *ConfigService) Page(pageInfo *view.ConfigPageView) (err error, res *common.PageInfo) {
	err, res = s.configDao.Page(pageInfo)
	if err != nil {
		return err, nil
	}
	return s.viewUtils.PageData2ViewList(res)
}

func (s *ConfigService) List(v *view.ConfigView) (err error, views []*view.ConfigView) {
	err, data := s.viewUtils.View2Data(v)
	if err != nil {
		return err, nil
	}
	var datas []*model2.Config
	if err, datas = s.configDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = s.viewUtils.Data2ViewList(datas)
		return
	}
}

// SelectConfigByKey 根据key查询Config记录
func (s *ConfigService) SelectConfigByKey(key string) (error, *view.ConfigView) {
	if err, config := s.configDao.SelectConfigByKey(key); err != nil {
		return err, nil
	} else {
		if err, configView := s.viewUtils.Data2View(config); err != nil {
			return err, nil
		} else {
			return nil, configView
		}
	}
}
