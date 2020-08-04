package router

import (
	"github.com/gin-gonic/gin"
	"goAdmin/ctrl"
	"goAdmin/handler"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UseRouter := Router.Group("user").
		Use(handler.JWTAuth()).
		Use(handler.CasbinHandler())
	{
		UseRouter.POST("ChangePassword", ctrl.ChangePassword) //修改密码
	}
}

//设置基础网关服务
func InitBaseRiyter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", ctrl.Register) //注册用户
		BaseRouter.POST("login", ctrl.Login)       //登录用户
		BaseRouter.POST("captcha", ctrl.Captcha)   //获取验证码
	}
}
