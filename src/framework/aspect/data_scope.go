package aspect

import (
	"fmt"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/common/constants"
	"strings"
)

const (
	// DATA_SCOPE_ALL 全部数据权限
	DATA_SCOPE_ALL = "1"
	// DATA_SCOPE_CUSTOM 自定数据权限
	DATA_SCOPE_CUSTOM = "2"
	// DATA_SCOPE_DEPT 部门数据权限
	DATA_SCOPE_DEPT = "3"
	// DATA_SCOPE_DEPT_AND_CHILD 部门及以下数据权限
	DATA_SCOPE_DEPT_AND_CHILD = "4"
	// DATA_SCOPE_SELF 仅本人数据权限
	DATA_SCOPE_SELF = "5"
)

// DataScopeFilter 数据权限过滤
func DataScopeFilter(user *vo.UserView, deptAlias string, userAlias string, permission string) string {
	// 如果是超级管理员，则不过滤数据
	if user == nil {
		return ""
	}
	if user.Id == constants.SYSTEM_ROLE_ADMIN_ID {
		return ""
	}
	sqlString := strings.Builder{}
	conditions := make([]string, 0)
	for _, role := range user.Roles {
		dataScope := role.DataScope
		if dataScope != DATA_SCOPE_CUSTOM && contains(conditions, dataScope) {
			continue
		}
		if permission != "" && role.Permissions != nil && !contains(role.Permissions, permission) {
			continue
		}
		switch dataScope {
		case DATA_SCOPE_ALL:
			sqlString.Reset()
			conditions = []string{dataScope}
			break
		case DATA_SCOPE_CUSTOM:
			sqlString.WriteString(fmt.Sprintf(" OR %s.id IN ( SELECT dept_id FROM sys_role_dept WHERE role_id = %s ) ", deptAlias, role.Id))
			break
		case DATA_SCOPE_DEPT:
			sqlString.WriteString(fmt.Sprintf(" OR %s.id = %s ", deptAlias, user.DeptId))
			break
		case DATA_SCOPE_DEPT_AND_CHILD:
			sqlString.WriteString(fmt.Sprintf(" OR %s.id IN ( SELECT id FROM sys_dept WHERE id = %s or find_in_set( %s , ancestors ) )", deptAlias, user.DeptId, user.DeptId))
			break
		case DATA_SCOPE_SELF:
			if userAlias != "" {
				sqlString.WriteString(fmt.Sprintf(" OR %s.id = %s ", userAlias, user.Id))
			} else {
				sqlString.WriteString(fmt.Sprintf(" OR %s.id = 0 ", deptAlias))
			}
			break
		}
		conditions = append(conditions, dataScope)
	}
	if len(conditions) == 0 {
		sqlString.WriteString(fmt.Sprintf(" OR %s.dept_id = 0 ", deptAlias))
	}
	if sqlString.Len() > 0 {
		return fmt.Sprintf("%s", sqlString.String()[4:])
	}
	return ""
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
