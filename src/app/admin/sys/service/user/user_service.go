// Package user 自动生成模板 UserService
// @description <TODO description class purpose>
// @author
// @File: user
// @version 1.0.0
// @create 2023-08-18 14:02:24
package user

import (
	"errors"
	"gorm.io/gorm"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/app/admin/sys/service/dept"
	"matuto-base/src/app/admin/sys/service/role"
	"matuto-base/src/common"
	"matuto-base/src/common/constants"
	"matuto-base/src/framework/aspect"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"matuto-base/src/utils/convert"
)

type Service struct {
	userDao     dao.UserDao
	deptService dept.DeptService
	roleService role.RoleService
	userRoleDao dao.UserRoleDao
	userPostDao dao.UserPostDao
}

// Create 创建User记录
// Author
func (s *Service) Create(userView *vo.UserView) (err error) {
	//err1, user := s.viewUtils.View2Data(userView)
	err1, user := convert.View2Data[vo.UserView, model.User](userView)
	if err1 != nil {
		return err1
	}
	tx := global.GormDao.Begin()
	if err = s.userDao.Create(tx, *user); err != nil {
		return err
	} else {
		if userView.RoleIds != nil && len(userView.RoleIds) > 0 {
			if err2 := s.insertUserRole(tx, user.Id, userView.RoleIds); err2 != nil {
				tx.Rollback()
				return err2
			}
		}
		if userView.PostIds != nil && len(userView.PostIds) > 0 {
			if err3 := s.insertUserPost(tx, user.Id, userView.PostIds); err3 != nil {
				tx.Rollback()
				return err3
			}
		}
		// 提交事务
		tx.Commit()
		return nil
	}
}

// insertUserPost 插入用户岗位关联数据
func (s *Service) insertUserPost(tx *gorm.DB, id string, ids []string) error {
	var userPosts []model.UserPost
	for _, postId := range ids {
		userPosts = append(userPosts, model.UserPost{
			UserId: id,
			PostId: postId,
		})
	}
	return s.userPostDao.CreateBatch(tx, userPosts)
}

// insertUserRole 插入用户角色关联数据
func (s *Service) insertUserRole(tx *gorm.DB, userId string, roleIds []string) error {
	var userRoles []model.UserRole
	for _, roleId := range roleIds {
		userRoles = append(userRoles, model.UserRole{
			UserId: userId,
			RoleId: roleId,
		})
	}
	return s.userRoleDao.CreateBatch(tx, userRoles)
}

// DeleteByIds 批量删除User记录
// Author
func (s *Service) DeleteByIds(ids []string, loginUserId string) (err error) {
	for _, id := range ids {
		if constants.SYSTEM_ADMIN_ID == id {
			return errors.New("不允许操作超级管理员用户")
		}
		if err = s.CheckUserDataScope(id, loginUserId); err != nil {
			return err
		}
	}
	tx := global.GormDao.Begin()
	// 删除用户角色关联数据
	if err = s.userRoleDao.DeleteByUserIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	// 删除用户岗位关联数据
	if err = s.userPostDao.DeleteByUserIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	err = s.userDao.DeleteByIds(tx, ids)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return err
}

// Update 更新User记录
// Author
func (s *Service) Update(id string, userView *vo.UserView) (err error) {
	userView.Id = id
	err1, user := convert.View2Data[vo.UserView, model.User](userView)
	if err1 != nil {
		return err1
	}
	err = s.userDao.Update(*user)
	return err
}

// Get 根据id获取User记录
// Author
func (s *Service) Get(id string) (err error, userView *vo.UserView) {
	if id == "" {
		return nil, nil
	}
	err1, user := s.userDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	if err, userView = convert.Data2View[vo.UserView, model.User](user); err != nil {
		return err, nil
	} else {
		if err2, deptView := s.deptService.Get(userView.DeptId); err2 != nil {
			return err2, nil
		} else {
			userView.Dept = deptView
		}
		// 组装角色信息
		if err3, roles := s.roleService.AssembleRolesByUserId(id); err3 != nil {
			return err3, nil
		} else {
			userView.Roles = roles
		}
		return
	}
}

// Page 分页获取User记录
// Author
func (s *Service) Page(pageInfo *vo.UserPageView, user *vo.UserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.userDao.Page(pageInfo); err != nil {
		return err, res
	}
	if err, views := convert.PageData2ViewList[vo.UserView, model.User](res); err != nil {
		//if err, res = s.viewUtils.PageData2ViewList(res); err != nil {
		return err, res
	} else {
		if o, ok := views.Rows.([]*vo.UserView); ok {
			// 组装部门数据
			for i := 0; i < len(o); i++ {
				deptId := o[i].DeptId
				if err3, deptView := s.deptService.Get(deptId); err3 != nil {
					return err3, nil
				} else {
					o[i].Dept = deptView
				}
			}
		}
		return err, views
	}
}

// List 获取User记录
func (s *Service) List(v *vo.UserView) (err error, views []*vo.UserView) {
	err, data := convert.View2Data[vo.UserView, model.User](v)
	if err != nil {
		return err, nil
	}
	var datas []*model.User
	if err, datas = s.userDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = convert.Data2ViewList[vo.UserView, model.User](datas)
		return
	}
}

// GetByUserName 根据userName获取User记录
// Author
func (s *Service) GetByUserName(userName string) (err error, userView *vo.UserView) {
	err1, user := s.userDao.GetByUserName(userName)
	if err1 != nil {
		return err1, nil
	}
	err2, userView := convert.Data2View[vo.UserView, model.User](user)
	if err2 != nil {
		return err2, nil
	}
	return
}

// CheckFieldUnique 校验字段是否唯一
// Author
func (s *Service) CheckFieldUnique(fieldName, value, id string) error {
	if fieldName == "" || value == "" {
		return nil
	}
	if err, data := s.userDao.SelectByField(fieldName, value); err != nil {
		return err
	} else {
		if data != nil && data.Id != id {
			return errors.New("数据重复")
		}
		return nil
	}
}

// CheckUserDataScope 校验数据权限
func (s *Service) CheckUserDataScope(userId, loginUserId string) error {
	if constants.SYSTEM_ADMIN_ID != loginUserId {
		err, userView := s.Get(userId)
		if err != nil {
			return err
		}
		// 数据权限控制
		// err, data := s.viewUtils.View2Data(userView)
		if err != nil {
			return err
		}
		filter := aspect.DataScopeFilter(userView, "d", "u", "")
		param := &model.User{}
		param.Id = userId
		param.DataScopeSql = filter
		// data.DataScopeSql = filter
		err, _ = s.userDao.List(param)
		if err != nil {
			return err
		}
	}
	return nil
}

// ResetPwd 重置密码
func (s *Service) ResetPwd(v *vo.UserView) error {
	err, user := convert.View2Data[vo.UserView, model.User](v)
	if err != nil {
		return err
	}
	salt := utils.GenUID()
	user.Password = utils.EncryptionPassword(user.Password, salt)
	user.Salt = salt
	return s.userDao.Update(*user)
}

// ChangeStatus 更新状态
func (s *Service) ChangeStatus(v *vo.UserView) error {
	err, user := convert.View2Data[vo.UserView, model.User](v)
	if err != nil {
		return err
	}
	return s.userDao.Update(*user)
}

// AuthRole	角色授权
func (s *Service) AuthRole(v *vo.UserView) error {
	tx := global.GormDao.Begin()
	// 删除用户角色关联数据
	if err := s.userRoleDao.DeleteByUserIds(tx, []string{v.Id}); err != nil {
		tx.Rollback()
		return err
	}
	// 插入用户角色关联数据
	if err := s.insertUserRole(tx, v.Id, v.RoleIds); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// SelectAllocatedList 获取已分配用户角色的用户列表
func (s *Service) SelectAllocatedList(pageInfo *vo.UserPageView, user *vo.UserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.userDao.SelectAllocatedList(pageInfo); err != nil {
		return err, nil
	}
	if err, res = convert.PageData2ViewList[model.User, vo.UserView](res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*vo.UserView); ok {
			// 组装部门数据
			for i := 0; i < len(o); i++ {
				deptId := o[i].DeptId
				if err3, deptView := s.deptService.Get(deptId); err3 != nil {
					return err3, nil
				} else {
					o[i].Dept = deptView
				}
			}
		}
		return err, res
	}
}

// SelectUnallocatedList 获取未分配用户角色的用户列表
func (s *Service) SelectUnallocatedList(pageInfo *vo.UserPageView, user *vo.UserView) (err error, res *common.PageInfo) {
	pageInfo.DataScopeSql = aspect.DataScopeFilter(user, "d", "u", "")
	if err, res = s.userDao.SelectUnallocatedList(pageInfo); err != nil {
		return err, nil
	}
	if err, res = convert.PageData2ViewList[model.User, vo.UserView](res); err != nil {
		return err, res
	} else {
		if o, ok := res.Rows.([]*vo.UserView); ok {
			// 组装部门数据
			for i := 0; i < len(o); i++ {
				deptId := o[i].DeptId
				if err3, deptView := s.deptService.Get(deptId); err3 != nil {
					return err3, nil
				} else {
					o[i].Dept = deptView
				}
			}
		}
		return err, res
	}
}

// GetByDeptId 根据部门id获取User记录
func (s *Service) GetByDeptId(deptId string) (err error, userView []*vo.UserView) {
	err1, user := s.userDao.GetByDeptId(deptId)
	if err1 != nil {
		return err1, nil
	}
	//err2, userView := s.viewUtils.Data2ViewList(user)
	err2, userView := convert.Data2ViewList[vo.UserView, model.User](user)
	if err2 != nil {
		return err2, nil
	}
	return
}
