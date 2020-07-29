package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID      uuid.UUID `json:"uuid" gorm:"comment:'用户UUID'"`
	LoginName string    `json:"loginName" gorm:"comment:'登录用户名'"`
	Name      string    `json:"name" gorm:"defalult:'系统用户';comment:'用户昵称'"`
	Image     string    `json:"image" gorm:"comment:'用户头像'"`
	RoleId    int       `json:"roleId" gorm:"comment:'用户角色id'"`
	Password  string    `json:"password" gorm:"comment:'用户密码'"`
}
