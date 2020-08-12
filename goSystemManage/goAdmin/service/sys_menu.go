package service

import (
	"errors"
	"goAdmin/global"
	"goAdmin/model"
)

func AllMenuList() (err error, menus []model.SysMenu) {
	var list []model.SysMenu
	//第一步，获取所有menu数据
	err = global.GLOBAL_DB.Order("sort", true).Find(&list).Error
	if err != nil {
		//第二步，将数据进行树形组装
		menus = getMenuTree(list)
		return err, menus
	} else {
		return
	}
}

func getMenuTree(list []model.SysMenu) (menus []model.SysMenu) {
	treeMap := make(map[int][]model.SysMenu)
	for _, v := range list {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	menus = treeMap[0] //获取所有根菜单
	for i := 0; i < len(menus); i++ {
		getChildrenMenuList(&menus[i], treeMap)
	}
}
func getChildrenMenuList(menu *model.SysMenu, treemap map[int][]model.SysMenu) {
	menu.Children = treemap[int(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		getChildrenMenuList(&menu.Children[i], treemap)
	}
}

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
