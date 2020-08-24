package handler

import (
	"github.com/gin-gonic/gin"
	"goAdmin/base/request"
	"goAdmin/global"
	"goAdmin/service"
)

//拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, flag := c.Get("claims")
		if flag {
			waitUse := claims.(*request.CustomClaims)
			//获取请求url
			obj := c.Request.URL.RequestURI()
			//获取请求办法
			act := c.Request.Method
			//获取用户角色
			sub := waitUse.RoleId
			e := service.Casbin()
			if global.GLOBAL_CONFIG.System.Env == "public" || e.Enforce(sub, obj, act) {
				c.Next()
			} else {
				c.Abort()
				return
			}
		} else {
			return
		}
	}
}
