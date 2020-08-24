package model

import "github.com/jinzhu/gorm"

//后续考虑将jwt的内容放到redis内
type SysJwtBlack struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:'jwt'"`
}
