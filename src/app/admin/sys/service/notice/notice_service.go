// Package notice 自动生成模板 NoticeService
// @description <TODO description class purpose>
// @author matuto
// @version 1.0.0
// @create 2023-09-12 13:45:22
package notice

import (
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/service/notice/view"
	"matuto-base/src/common"
)

type NoticeService struct {
	noticeDao dao.NoticeDao
	viewUtils view.NoticeViewUtils
}

// Create 创建通知公告表记录
// Author matuto
func (s *NoticeService) Create(notice *view.NoticeCreateView) error {
	return s.noticeDao.Create(notice)
}

// DeleteByIds 批量删除通知公告表记录
// Author matuto
func (s *NoticeService) DeleteByIds(ids []string) error {
	return s.noticeDao.DeleteByIds(ids)
}

// Update 更新通知公告表记录
// Author matuto
func (s *NoticeService) Update(notice *view.NoticeEditView) error {
	return s.noticeDao.Update(notice)
}

// Get 根据id获取通知公告表记录
// Author matuto
func (s *NoticeService) Get(id string) (error, *view.NoticeView) {
	if err, notice := s.noticeDao.Get(id); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2View(notice)
	}
}

// Page 分页获取通知公告表记录
// Author matuto
func (s *NoticeService) Page(pageInfo *view.NoticePageView) (error, *common.PageInfo) {
	if err, res := s.noticeDao.Page(pageInfo); err != nil {
		return err, nil
	} else {
		return s.viewUtils.PageData2ViewList(res)
	}
}

// List 获取通知公告表列表
// Author matuto
func (s *NoticeService) List(v *view.NoticeQueryView) (error, []*view.NoticeView) {
	if err, dataList := s.noticeDao.List(v); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}
