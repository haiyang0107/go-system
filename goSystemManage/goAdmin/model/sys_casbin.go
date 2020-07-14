package model

type CasbinModel struct {
	ID          uint   `json:"id" gorm:"column:id;comment:'主键id'"`
	Ptype       string `json:"ptype" gorm:"column:ptype;comment:'类型'"`
	AuthorityId string `json:"roleName" gorm:"column:role_id;comment:'角色id'"`
	Path        string `json:"path" gorm:"column:path;comment:'请求路径'"`
	Method      string `json:"method" gorm:"column:method;comment:'请求方式'"`
}
