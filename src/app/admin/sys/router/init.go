package router

import (
	"github.com/gin-gonic/gin"
)

type SysRouter struct {
	userRouter     UserRouter
	menuRouter     MenuRouter
	postRouter     PostRouter
	configRouter   ConfigRouter
	deptRouter     DeptRouter
	roleRouter     RoleRouter
	dictTypeRouter DictTypeRouter
	dictDataRouter DictDataRouter
	noticeRouter   NoticeRouter
	monitorRouter  MonitorRouter
}

// var jobRouter job.JobRouter
// var jobLogRouter jobLog.JobLogRouter
// var logininForRouter logininfor.LogininforRouter

// InitSysRouter 初始化 Init 路由信息
func (r *SysRouter) InitSysRouter(Router *gin.RouterGroup) {
	sys := Router.Group("sys")
	{
		r.userRouter.InitUserRouter(sys)
		r.menuRouter.InitMenuRouter(sys)
		r.roleRouter.InitRoleRouter(sys)
		r.dictTypeRouter.InitDictTypeRouter(sys)
		r.dictDataRouter.InitDictDataRouter(sys)
		r.noticeRouter.InitNoticeRouter(sys)
		r.postRouter.InitPostRouter(sys)
		r.configRouter.InitConfigRouter(sys)
		r.deptRouter.InitDeptRouter(sys)
		// 监控相关
		r.monitorRouter.InitMonitorRouter(sys)
		//jobRouter.InitJobRouter(sys)
		//jobLogRouter.InitJobLogRouter(sys)
		//logininForRouter.InitLogininforRouter(sys)
	}
}
