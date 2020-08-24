package response

import "goAdmin/model"

type LoginResponse struct {
	User      model.SysUser `json:"user""`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"` //过期时间
}
