package response

type SysCaptcha struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}
