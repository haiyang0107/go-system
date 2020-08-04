package initialization

import (
	"github.com/gin-gonic/gin"
	"goAdmin/global"
	"goAdmin/handler"
	"goAdmin/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	//跨域处理
	Router.Use(handler.Cors())
	//添加路由分组，上线统一使用
	ApiGroup := Router.Group("")
	router.InitBaseRiyter(ApiGroup) //注册基础服务
	router.InitUserRouter(ApiGroup) //注册用户服务
	global.GLOBAL_LOG.Info("register routers for server Success ")
	return Router
}
