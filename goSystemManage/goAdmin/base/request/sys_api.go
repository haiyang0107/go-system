package request

import "goAdmin/model"

type ApiParams struct {
	model.SysApi
	PageStrut
	OrderString string `json:"orderString"`
	Desc        bool   `json:"desc"`
}
