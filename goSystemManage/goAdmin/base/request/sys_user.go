package request

//注册用户结构体
type RegisterStruct struct {
	LoginName string `json:"loginName"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Password  string `json:"password"`
}

type LoginStruct struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Captcha   string `json:captcha`
	CaptchaId string `json:captchaId`
}

type ChangePasswordStruct struct {
	LoginName   string `json:"loginName"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
