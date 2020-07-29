package ctrl

import (
	"github.com/gin-gonic/gin"
	. "github.com/mojocn/base64Captcha"
	"goAdmin/base/request"
	response2 "goAdmin/base/response"
	"goAdmin/global"
	"goAdmin/global/response"
)

//用户注册接口
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)

}

//用户登录接口
func Login(c *gin.Context) {

}

//验证码生成策略
func Captcha(c *gin.Context) {
	//字符生成策略
	driver := NewDriverDigit(global.GLOBAL_CONFIG.Captcha.ImgHeight, global.GLOBAL_CONFIG.Captcha.ImgWidth, global.GLOBAL_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := NewCaptcha(driver, DefaultMemStore)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMessage("验证码获取失败！", c)
	} else {
		response.Result(response.SUCCESS, response2.SysCaptcha{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
