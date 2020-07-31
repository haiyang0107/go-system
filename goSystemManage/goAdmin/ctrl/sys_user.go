package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goAdmin/base/request"
	response2 "goAdmin/base/response"
	"goAdmin/global"
	"goAdmin/global/response"
	"goAdmin/model"
	"goAdmin/service"
	"goAdmin/util"
)

//用户注册接口
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	userVerify := util.Rules{
		"LoginName": {util.NotEmpty()},
		"Name":      {util.NotEmpty()},
		"Password":  {util.NotEmpty()},
	}
	userVerifyErr := util.Verify(R, userVerify)
	if userVerifyErr != nil {
		response.FailWithMessage(userVerifyErr.Error(), c)
		return
	}
	user := &model.SysUser{Name: R.Name, LoginName: R.LoginName, Password: R.Password, Image: R.Image}
	err, userReturn := service.Register(*user)
	if err != nil {
		response.FailWithMsgAndData(err.Error(), userReturn, c)
	} else {
		response.SuccessWithMsgAndData("注册成功", userReturn, c)
	}

}

//用户登录接口
func Login(c *gin.Context) {

}

//验证码生成策略
func Captcha(c *gin.Context) {
	//字符生成策略
	driver := base64Captcha.NewDriverDigit(global.GLOBAL_CONFIG.Captcha.ImgHeight, global.GLOBAL_CONFIG.Captcha.ImgWidth, global.GLOBAL_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
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
