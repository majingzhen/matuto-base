// Package dept 自动生成模板 DeptService
// @description <TODO description class purpose>
// @author
// @File: dept
// @version 1.0.0
// @create 2023-08-21 10:27:01
package dept

import (
	"errors"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/dao"
	"matuto-base/src/app/admin/sys/model"
	"matuto-base/src/app/admin/sys/service/role"
	"matuto-base/src/app/admin/sys/service/user/extend"
	"matuto-base/src/common/constants"
	"matuto-base/src/framework/aspect"
	"matuto-base/src/global"
	"matuto-base/src/utils/convert"
	"strings"

	"gorm.io/gorm"
)

type DeptService struct {
	deptDao     dao.DeptDao
	userService extend.UserExtendService
	roleService role.RoleService
}

// Create 创建Dept记录
// Author
func (s *DeptService) Create(view *vo.DeptView) (err error) {
	// 判断名称是否重复
	err, exist := s.deptDao.CheckDeptNameALL(view.DeptName, view.ParentId)
	if exist {
		return errors.New("部门名称已经存在")
	}
	// 根据前端传入的数据创建model
	if err1, dept := convert.View2Data[vo.DeptView, model.Dept](view); err1 != nil {
		return err1
	} else {
		if err1, deptView := s.Get(dept.ParentId); err1 != nil {
			return errors.New("父级部门不存在")
		} else {
			if deptView.Status == constants.DEPT_DISABLE {
				return errors.New("父级部门已经被禁用, 禁止新增")
			}
			dept.Ancestors = deptView.Ancestors + "," + dept.ParentId
		}
		return s.deptDao.Create(*dept)
	}
}

// Delete 删除Dept记录
// Author
func (s *DeptService) Delete(id string) error {
	// 判断是否存在子集
	err, children := s.deptDao.SelectChildrenDeptById(id)
	if err != nil {
		return err
	} else {
		if len(*children) > 0 {
			return errors.New("存在子部门, 不允许删除")
		}
	}
	if err, userView := s.userService.GetByDeptId(id); err != nil {
		return err
	} else {
		if userView != nil {
			return errors.New("部门存在用户, 不允许删除")
		}
	}
	err = s.deptDao.Delete(id)
	return err
}

// Update 更新Dept记录
// Author
func (s *DeptService) Update(id string, deptView *vo.DeptView) (err error) {
	var oldDept *vo.DeptView
	// 判断是否存在
	if err, oldDept = s.Get(id); err != nil {
		return errors.New("部门不存在")
	}
	// 判断名称是否重复
	err, exist := s.deptDao.CheckDeptNameALL(deptView.DeptName, deptView.ParentId)
	if exist {
		return errors.New("部门名称已经存在")
	}
	if id == deptView.ParentId {
		return errors.New("上级部门不能是自己")
	}
	deptView.Id = id
	if err1, dept := convert.View2Data[vo.DeptView, model.Dept](deptView); err1 != nil {
		return err1
	} else {
		if err1, newParentDept := s.Get(dept.ParentId); err1 != nil {
			return errors.New("父级部门不存在")
		} else {
			newAncestors := newParentDept.Ancestors + "," + newParentDept.Id
			oldAncestors := oldDept.Ancestors
			dept.Ancestors = newAncestors
			// 更新子部门的祖级列表
			tx := global.GormDao.Begin()
			err = s.updateDeptChildren(tx, dept.Id, newAncestors, oldAncestors)
			if err != nil {
				tx.Rollback()
				return errors.New("数据更新失败")
			} else {
				if err = s.deptDao.Update(tx, *dept); err != nil {
					tx.Rollback()
					return err
				} else {
					tx.Commit()
					return nil
				}
			}
		}
	}
}

// updateDeptChildren 更新子部门的祖级列表
func (s *DeptService) updateDeptChildren(tx *gorm.DB, id, newAncestors, oldAncestors string) (err error) {
	var children *[]model.Dept
	err, children = s.deptDao.SelectChildrenDeptById(id)
	if err != nil {
		return err
	} else {
		for _, child := range *children {
			child.Ancestors = strings.Replace(child.Ancestors, oldAncestors, newAncestors, 1)
			err = s.deptDao.Update(tx, child)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Get 根据id获取Dept记录
// Author
func (s *DeptService) Get(id string) (err error, deptView *vo.DeptView) {
	if id == "" {
		return nil, nil
	}
	err1, dept := s.deptDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, deptView = convert.Data2View[vo.DeptView, model.Dept](dept)
	return err, deptView
}

// List 获取Dept列表
func (s *DeptService) List(v *vo.DeptView, userView *vo.UserView) (err error, views []*vo.DeptView) {
	err, data := convert.View2Data[vo.DeptView, model.Dept](v)
	if err != nil {
		return err, nil
	}
	data.DataScopeSql = aspect.DataScopeFilter(userView, "d", "u", "")
	var datas []*model.Dept
	if err, datas = s.deptDao.List(data); err != nil {
		return err, nil
	} else {
		err, views = convert.Data2ViewList[vo.DeptView, model.Dept](datas)
		return
	}
}

// SelectDeptTree 获取部门树
func (s *DeptService) SelectDeptTree(v *vo.DeptView, sysUserView *vo.UserView) (error, []*vo.DeptTreeView) {
	err, data := convert.View2Data[vo.DeptView, model.Dept](v)
	if err != nil {
		return err, nil
	}
	data.DataScopeSql = aspect.DataScopeFilter(sysUserView, "d", "u", "")
	var datas []*model.Dept
	if err, datas = s.deptDao.List(data); err != nil {
		return err, nil
	} else {
		var trees []*vo.DeptTreeView
		if err, trees = convert.Data2ViewList[vo.DeptTreeView, model.Dept](datas); err != nil {
			return err, nil
		} else {
			deptTrees := getDeptTree(trees)
			return nil, deptTrees
		}
	}
}

// SelectDeptTreeByRole 获取部门树（排除下级）
func (s *DeptService) SelectDeptTreeByRole(id string) (error, []string) {
	err, roleView := s.roleService.Get(id)
	if err != nil {
		return err, nil
	}
	return s.deptDao.SelectDeptListByRoleId(id, roleView.DeptCheckStrictly)
}

// getDeptTree 获取部门树
func getDeptTree(departments []*vo.DeptTreeView) []*vo.DeptTreeView {
	var rootDepts []*vo.DeptTreeView

	// 遍历所有部门，找到根节点
	for _, dept := range departments {
		if dept.ParentId == "0" {
			rootDepts = append(rootDepts, dept)
		}
	}
	// 递归获取每个根节点的子部门
	for i := range rootDepts {
		rootDepts[i].Children = getChildren(rootDepts[i].Id, departments)
	}
	return rootDepts
}

// getChildren 获取所有子集部门
func getChildren(parentId string, departments []*vo.DeptTreeView) []*vo.DeptTreeView {
	var children []*vo.DeptTreeView
	// 遍历所有部门，找到指定父节点的子部门
	for _, dept := range departments {
		if dept.ParentId == parentId {
			// 递归获取子部门的子部门
			dept.Children = getChildren(dept.Id, departments)
			children = append(children, dept)
		}
	}
	return children
}
