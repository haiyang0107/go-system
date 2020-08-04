package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"goAdmin/global"
	"goAdmin/model"
	"goAdmin/util"
)

//进行用户注册
func Register(user model.SysUser) (err error, userInfo model.SysUser) {
	db := global.GLOBAL_DB
	//判断用户是否注册
	flag := db.Where("name = ?", user.Name).First(&user).RecordNotFound()
	if !flag {
		return errors.New("用户已注册！"), userInfo
	} else {
		//否则，对密码进行加密，然后数据入库
		user.Password = util.Md5Enc([]byte(user.Password))
		user.UUID = uuid.NewV4()
		err := db.Create(&user).Error
		return err, user
	}
}

//用户登录 service 方法
func Login(user *model.SysUser) (err error, userInfo *model.SysUser) {
	var bean model.SysUser
	db := global.GLOBAL_DB
	util.Md5Enc([]byte(user.Password))
	err = db.Where("loginName = ? And password = ?", user.LoginName, user.Password).First(&bean).Error
	return err, &bean
}

//修改用户密码
func ChangePassword(user *model.SysUser, newPassword string) (err error, userInfo *model.SysUser) {
	var bean model.SysUser
	db := global.GLOBAL_DB
	util.Md5Enc([]byte(user.Password))
	err = db.Where("loginName = ? And password = ?", user.LoginName, user.Password).First(&bean).Update("password", util.Md5Enc([]byte(newPassword))).Error
	return err, user
}
