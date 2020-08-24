package model

import "github.com/jinzhu/gorm"

type SysRole struct {
	gorm.Model
	RoleId      int       `json:"roleId" gorm:"comment:'角色ID'"`
	RoleName    string    `json:"roleName" gorm:"comment:'角色名称'"`
	ParentId    int       `json:"parentId" gorm:"comment:'父角色ID'"`
	Children    []SysRole `json:"children"` //添加角色组，按照父子关系定义角色关联
	SysMenus    []SysMenu `json:"menus" gorm:"many2many:sys_role_menus"`
	SysDataRole []SysRole `json:"dataRoleId" gorm:"many2many:sys_data_role;association_jointable_foreignkey:data_role_id"` //添加数据权限关联主键，自动建表
}
