package model

import "github.com/jinzhu/gorm"

type SysMenu struct {
	gorm.Model
	Level     int       `json:"level" gorm:"comment:'菜单级别'"`
	ParentId  int       `json:"parentId" gorm:"comment:'父菜单ID'"`
	Name      string    `json:"name" gorm:"comment:'菜单名称'"`
	Desc      string    `json:"desc" gorm:"comment:'菜单描述'"`
	Sort      int       `json:"sort" gorm:"comment:'排序'"`
	Path      string    `json:"path" gorm:"comment:'路由path'"`
	Icon      string    `json:"icon" gorm:"comment:'菜单图标'"`
	Component string    `json:"component" gorm:"comment:'对应前端路径'"`
	Children  []SysMenu `json:"children"`
	SysRoles  []SysRole `json:"sysRoles" gorm:"many2many:sys_role_menus"`
}
