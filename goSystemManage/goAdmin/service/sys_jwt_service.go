package service

import (
	"goAdmin/global"
	"goAdmin/model"
)

func IsBlackList(jwt string, jwtList model.SysJwt) bool {
	isNotFound := global.GLOBAL_DB.Where("jwt = ?", jwt).First(&jwtList).RecordNotFound()
	return !isNotFound
}