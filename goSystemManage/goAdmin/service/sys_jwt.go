package service

import (
	"goAdmin/global"
	"goAdmin/model"
)

//判断 jwt 是否是在黑名单中
func IsBlackList(jwt string, jwtList model.SysJwtBlack) bool {
	isNotFound := global.GLOBAL_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}

//将jwt 加入黑名单
func JoinInBlackList(jwtList model.SysJwtBlack) (err error) {
	err = global.GLOBAL_DB.Create(&jwtList).Error
	return err
}

//将jwt 存入redis 中
func SetRedisJwt(jwt model.SysJwtBlack, userName string) (err error) {
	//超时时间 一天
	err = global.GLOBAL_REDIS.Set(userName, jwt.Jwt, 1000*1000*1000*60*60*24).Err()
	return err
}

//从redis中取出 jwt 信息
func GetRedisJwt(userName string) (err error, redisJwt string) {
	redisJwt, err = global.GLOBAL_REDIS.Get(userName).Result()
	return err, redisJwt
}
