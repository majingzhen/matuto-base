package basic

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/sys/api/vo"
	"matuto-base/src/app/admin/sys/service/user"
	"matuto-base/src/global"
)

type BasicApi struct {
	userService user.Service
}

// GetLoginUser 获取当前登录用户
func (api *BasicApi) GetLoginUser(c *gin.Context) *vo.UserView {
	err, view := api.userService.Get(api.GetLoginUserId(c))
	if err != nil {
		global.Logger.Error("[获取登录用户] is error", zap.Error(err))
		return nil
	}
	return view
}

// GetLoginUserId 获取当前登录用户id
func (api *BasicApi) GetLoginUserId(c *gin.Context) string {
	userId := c.GetString("user_id")
	return userId
}

// GetLoginUserName 获取当前登录用户名
func (api *BasicApi) GetLoginUserName(c *gin.Context) string {
	userName := c.GetString("user_name")
	return userName
}
