package model

type SysRoleMenu struct {
	SysMenu
	MenuId   int           `json:"menuId" gorm:"comment:'菜单id'"`
	RoleId   int           `json:"roleId" gorm:"comment:'角色id'"`
	Children []SysRoleMenu `json:"children"`
}
