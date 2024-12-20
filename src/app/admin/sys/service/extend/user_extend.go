package extend

import (
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/utils/convert"
)

type UserExtendService struct {
	userDao dao.UserDao
}

// GetByDeptId 根据部门id获取User记录
func (s *UserExtendService) GetByDeptId(deptId string) (err error, userView []*vo.UserView) {
	err1, user := s.userDao.GetByDeptId(deptId)
	if err1 != nil {
		return err1, nil
	}
	err2, userView := convert.Data2ViewList[vo.UserView, model.User](user)
	if err2 != nil {
		return err2, nil
	}
	return
}
