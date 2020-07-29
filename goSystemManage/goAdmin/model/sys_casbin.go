package model

import "github.com/jinzhu/gorm"

type SysCasbin struct {
	gorm.Model
	Type   string `json:"type" gorm:"column:type;comment:'类型'"`
	RoleId string `json:"roleId" gorm:"column:role_id;comment:'角色id'"`
	Path   string `json:"path" gorm:"column:path;comment:'请求路径'"`
	Method string `json:"method" gorm:"column:method;comment:'请求方式'"`
}
