package handler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goAdmin/base/request"
	"goAdmin/global"
	"goAdmin/global/response"
	"goAdmin/model"
	"goAdmin/service"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpried     = errors.New("Token is expried")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

//产生新的jwt
func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(global.GLOBAL_CONFIG.Jwt.SigningKey)}
}

//创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.SigningKey)
}

//解析token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpried
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	} else {
		return nil, TokenInvalid
	}
}

//更新token
func (j *JWT) RefreshToken(tokenstring string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenstring, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", nil
	}
	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

//jwt 身份信息
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//我们这里jwt鉴权头部信息：x-token 登录时返回 token信息 前端进行token信息存储
		//记录过期时间，约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		modelToken := model.JwtBlacklist{
			Jwt: token,
		}
		if token == "" {
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if service.IsBlackList(token, modelToken) {
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, "你的令牌异地登录或者令牌失效", c)
			c.Abort()
			return
		}
		j := NewJWT()
		//解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpried {
				response.Result(response.ERROR, gin.H{
					"reload": true,
				}, "授权已过期", c)
				c.Abort()
				return
			}
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
