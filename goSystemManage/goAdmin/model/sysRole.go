package model

import "github.com/jinzhu/gorm"

type SysRole struct {
	gorm.Model
	RoleId   int    `json:"roleId" gorm:"comment:'角色ID'"`
	RoleName string `json:"roleName" gorm:"comment:'角色名称'"`
	ParentId int    `json:"parentId" gorm:"comment:'父角色ID'"`
}
