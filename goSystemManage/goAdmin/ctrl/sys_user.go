package ctrl

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goAdmin/base/request"
	resp "goAdmin/base/response"
	"goAdmin/global"
	"goAdmin/global/response"
	"goAdmin/handler"
	"goAdmin/model"
	"goAdmin/service"
	"goAdmin/util"
	"time"
)

//定义全局验证码工具，为登录做校验
var defaultMemStore = base64Captcha.DefaultMemStore

//用户注册接口
func Register(c *gin.Context) {
	var R request.RegisterStruct
	err := c.ShouldBindJSON(&R)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}
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

//用户---修改密码
func ChangePassword(c *gin.Context) {
	var ChangeEntity request.ChangePasswordStruct
	err := c.ShouldBindJSON(&ChangeEntity)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}
	userVerify := util.Rules{
		"LoginName":   {util.NotEmpty()},
		"OldPassword": {util.NotEmpty()},
		"NewPassword": {util.NotEmpty()},
	}
	userVerifyErr := util.Verify(ChangeEntity, userVerify)
	if userVerifyErr != nil {
		response.FailWithMessage(userVerifyErr.Error(), c)
	}
	u := &model.SysUser{LoginName: ChangeEntity.LoginName, Password: ChangeEntity.OldPassword}
	if err, _ := service.ChangePassword(u, ChangeEntity.NewPassword); err != nil {
		response.FailWithMessage("修改密码失败，请检查密码是否正确", c)
	} else {
		response.SuccessWithMessage("修改密码成功", c)
	}
}

//用户登录接口
func Login(c *gin.Context) {
	var Login request.LoginStruct
	err := c.ShouldBindJSON(&Login)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}
	userVerify := util.Rules{
		"LoginName": {util.NotEmpty()},
		"Captcha":   {util.NotEmpty()},
		"CaptchaId": {util.NotEmpty()},
		"Password":  {util.NotEmpty()},
	}
	userVerifyErr := util.Verify(Login, userVerify)
	if userVerifyErr != nil {
		response.FailWithMessage(userVerifyErr.Error(), c)
	}
	//校验验证码是否正确
	if defaultMemStore.Verify(Login.CaptchaId, Login.Captcha, true) {
		err, userBean := service.Login(&model.SysUser{LoginName: Login.LoginName, Password: Login.Password})
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		//生成token
		tokenJwt(c, *userBean)
	} else {
		response.FailWithMessage("验证码错误！", c)
	}
}

//用户管理--删除用户
func DeleteUser(c *gin.Context) {
	var ID request.GetById
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}

	IdVerifyErr := util.Verify(ID, util.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
	}
	if err := service.DeleteUser(ID.Id); err != nil {
		response.FailWithMessage("删除用户失败，请检查Id是否正确", c)
	} else {
		response.SuccessWithMessage("删除用户成功", c)
	}
}

// 用户管理-新增用户
func CreateUser(c *gin.Context) {
	var user model.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage("数据格式异常", c)
	}
	userVerify := util.Rules{
		"LoginName": {util.NotEmpty()},
		"Name":      {util.NotEmpty()},
		"Password":  {util.NotEmpty()},
		"RoleId":    {util.NotEmpty()},
		//"Image":  {util.NotEmpty()},后续添加 图片上传功能
	}
	userVerifyErr := util.Verify(user, userVerify)
	if userVerifyErr != nil {
		response.FailWithMessage(userVerifyErr.Error(), c)
		return
	}
	e, _ := service.Register(user)
	if e != nil {
		response.FailWithMessage(fmt.Sprintf("新增用户失败，原因是：%v", e), c)
	} else {
		response.SuccessWithMessage("新增用户成功", c)
	}
}

//登录之后，生成jwt
func tokenJwt(c *gin.Context, user model.SysUser) {
	j := &handler.JWT{
		SigningKey: []byte(global.GLOBAL_CONFIG.Jwt.SigningKey),
	}
	claims := request.CustomClaims{
		UUID:   user.UUID,
		ID:     user.ID,
		Name:   user.LoginName,
		RoleId: user.RoleId,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: time.Now().Unix() + 60*60*24*7,      // 过期时间 一周
			Issuer:    global.GLOBAL_CONFIG.Jwt.SigningKey, // 签名的发行者
			NotBefore: time.Now().Unix() - 1000,            // 签名生效时间
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败：", c)
		return
	}
	if global.GLOBAL_CONFIG.System.AllowSingle {
		response.SuccessWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.ExpiresAt * 1000,
		}, c)
		return
	}
	//加入Redis 管理jwt
}

//验证码生成策略
func Captcha(c *gin.Context) {
	//字符生成策略
	driver := base64Captcha.NewDriverDigit(global.GLOBAL_CONFIG.Captcha.ImgHeight, global.GLOBAL_CONFIG.Captcha.ImgWidth, global.GLOBAL_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, defaultMemStore)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMessage("验证码获取失败！", c)
	} else {
		response.Result(response.SUCCESS, resp.SysCaptcha{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
