package router

import (
	"github.com/gin-gonic/gin"
	"goAdmin/ctrl"
	"goAdmin/handler"
)

//设置用户基础路由
func InitUserRouter(Router *gin.RouterGroup) {
	UseRouter := Router.Group("user").
		Use(handler.JWTAuth()).
		Use(handler.CasbinHandler())
	{
		UseRouter.POST("ChangePassword", ctrl.ChangePassword) //修改密码
		UseRouter.POST("DeleteUser", ctrl.DeleteUser)         //删除用户
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

//设置菜单基础路由
func InitMenuRouter(Router *gin.RouterGroup) {
	UseRouter := Router.Group("menu").
		Use(handler.JWTAuth()).
		Use(handler.CasbinHandler())
	{
		UseRouter.POST("GetMenuPage", ctrl.GetMenuPage) //获取菜单列表

	}
}

//设置api请求路由
func InitApiRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api").
		Use(handler.JWTAuth()).
		Use(handler.CasbinHandler())
	{
		ApiRouter.POST("CreateApi", ctrl.CreateApi)     //创建api
		ApiRouter.POST("UpdateApi", ctrl.UpdateApi)     //更新api
		ApiRouter.POST("DeleteApi", ctrl.DeleteApi)     //删除api
		ApiRouter.POST("GetById", ctrl.GetById)         //根据id获取api详情
		ApiRouter.POST("AllListApi", ctrl.AllListApi)   //获取全部API，不分页
		ApiRouter.POST("PageListApi", ctrl.PageListApi) //分页api list 信息
	}
}
