// Package notice 自动生成模板 NoticeService
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-12 13:45:22
package service

import (
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/common"
	"matuto-base/src/utils/convert"
)

type NoticeService struct {
	noticeDao dao.NoticeDao
}

// Create 创建通知公告表记录
// Author matuto
func (s *NoticeService) Create(notice *vo.NoticeCreateView) error {
	err, noticeData := convert.View2Data[vo.NoticeCreateView, model.Notice](notice)
	if err != nil {
		return err
	}
	return s.noticeDao.Create(noticeData)
}

// DeleteByIds 批量删除通知公告表记录
// Author matuto
func (s *NoticeService) DeleteByIds(ids []string) error {
	return s.noticeDao.DeleteByIds(ids)
}

// Update 更新通知公告表记录
// Author matuto
func (s *NoticeService) Update(notice *vo.NoticeEditView) error {
	err, noticeData := convert.View2Data[vo.NoticeEditView, model.Notice](notice)
	if err != nil {
		return err
	}
	return s.noticeDao.Update(noticeData)
}

// Get 根据id获取通知公告表记录
// Author matuto
func (s *NoticeService) Get(id string) (error, *vo.NoticeView) {
	if err, notice := s.noticeDao.Get(id); err != nil {
		return err, nil
	} else {
		return convert.Data2View[vo.NoticeView, model.Notice](notice)
	}
}

// Page 分页获取通知公告表记录
// Author matuto
func (s *NoticeService) Page(pageInfo *vo.NoticePageView) (error, *common.PageInfo) {
	if err, res := s.noticeDao.Page(pageInfo); err != nil {
		return err, nil
	} else {
		return convert.PageData2ViewList[vo.NoticeView, model.Notice](res)
	}
}

// List 获取通知公告表列表
// Author matuto
func (s *NoticeService) List(v *vo.NoticeQueryView) (error, []*vo.NoticeView) {
	if err, dataList := s.noticeDao.List(v); err != nil {
		return err, nil
	} else {
		return convert.Data2ViewList[vo.NoticeView, model.Notice](dataList)
	}
}
