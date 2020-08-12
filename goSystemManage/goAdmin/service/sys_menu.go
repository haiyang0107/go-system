package service

import (
	"errors"
	"goAdmin/global"
	"goAdmin/model"
)

func GetMenuInfoById(id int) (err error, menu model.SysMenu) {
	err = global.GLOBAL_DB.Where(" id = ? ", id).First(&menu).Error
	return
}

func CreatMenu(menu model.SysMenu) (err error) {
	err = global.GLOBAL_DB.Where("name = ? ", menu.Name).First(&model.SysMenu{}).Error
	if err != nil {
		return errors.New("菜单名称存在重复值,请修改name")
	} else {
		return global.GLOBAL_DB.Create(menu).Error
	}
}

func UpdateMenu(menu model.SysMenu) (err error) {
	var oldMenu model.SysMenu
	db := global.GLOBAL_DB.Where(" id = ? ", menu.ID).Find(&oldMenu)
	if oldMenu.Name != menu.Name {
		//需要判断是否存在重复数据
		flag := global.GLOBAL_DB.Where("id <> ? AND name = ? ", menu.ID, menu.Name).First(&model.SysMenu{}).RecordNotFound()
		if !flag {
			return errors.New("存在相同的菜单name,修改失败！")
		}
	}
	return db.Updates(menu).Error
}

func DeleteMenuById(id int) (e error) {
	err := global.GLOBAL_DB.Where(" parent_id = ? ", id).First(&model.SysMenu{}).Error
	if err != nil {

	} else {
		return errors.New("此菜单存在子菜单不可删除！")
	}
}
