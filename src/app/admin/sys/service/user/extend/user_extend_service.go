package extend

import (
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/service/user/view"
)

type UserExtendService struct {
	userDao   dao.UserDao
	viewUtils view.UserViewUtils
}

// GetByDeptId 根据部门id获取User记录
func (s *UserExtendService) GetByDeptId(deptId string) (err error, userView []*view.UserView) {
	err1, user := s.userDao.GetByDeptId(deptId)
	if err1 != nil {
		return err1, nil
	}
	err2, userView := s.viewUtils.Data2ViewList(user)
	if err2 != nil {
		return err2, nil
	}
	return
}
