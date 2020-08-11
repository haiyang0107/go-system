package service

import (
	"goAdmin/base/request"
	"goAdmin/global"
	"goAdmin/model"
)

func GetMenuPage(page request.PageStrut) (err error, list interface{}, total int) {
	var menuList []model.SysMenu
	db := global.GLOBAL_DB
	var count int32
	db.Model(model.SysMenu{}).Where("parentId = ? ", -1).Count(&count)
	err = db.Model(&model.SysMenu{}).Where("parentId = ? ", -1).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&menuList).Error
	return err, menuList, total
}
