package service

import (
	"errors"
	"goAdmin/base/request"
	"goAdmin/global"
	"goAdmin/model"
)

//创建api接口
func CreateApi(api model.SysApi) (err error) {
	//api需要判重
	db := global.GLOBAL_DB
	info := db.Where("path = ? AND method = ?", api.Path, api.Method).Find(&model.SysApi{}).Error
	if info == nil {
		return errors.New("存在相同数据")
	} else {
		return db.Create(&api).Error
	}
}

//修改api
func UpdateApi(api model.SysApi) (err error) {
	var a model.SysApi
	db := global.GLOBAL_DB
	error := db.Where("id = ? ", api.ID).First(&a).Error
	if error != nil {
		return error
	} else {
		if a.Method != api.Method || a.Path != api.Path {
			//判断一下新修改的数据是否存在重复数据
			info := db.Where("path = ? AND method = ?", api.Path, api.Method).Find(&model.SysApi{}).RecordNotFound()
			if !info {
				return errors.New("存在相同api数据")
			} else {
				//后期需要更新权限信息
				err = UpdateCasbinApi(a.Path, a.Method, api.Path, api.Method)
				if err != nil {
					return err
				} else {
					err = db.Save(&api).Error
					return err
				}
			}
		}
	}
}

func DeleteApi(id int) (err error) {
	db := global.GLOBAL_DB
	var api model.SysApi
	flag := db.Where("id = ? ", id).First(&api).RecordNotFound()
	if !flag {
		return errors.New("数据找不到异常")
	} else {
		err = db.Delete(api).Error
		clearCabinApi(1, api.Method, api.Path)
		return err
	}
}

func GetApiById(id int) (err error, m model.SysApi) {
	err = global.GLOBAL_DB.Where("id = ? ", id).First(&model.SysApi{}).Error
	return
}

func GetAllListApi() (err error, apis []model.SysApi) {
	err = global.GLOBAL_DB.Find(&apis).Error
	return
}

func PageApisList(api model.SysApi, page request.PageStrut, desc bool, orderString string) (err error, list []model.SysApi, total int) {
	db := global.GLOBAL_DB.Model(&model.SysApi{})
	limit := page.PageSize
	offset := page.PageSize * (page.PageNum - 1)
	var apiList []model.SysApi

	if api.Path != "" {
		db.Where("path LIKE ? ", "%"+api.Path+"%")
	}
	if api.Description != "" {
		db.Where("description LIKE ? ", "%"+api.Description+"%")
	}
	if api.Method != "" {
		db.Where("method = ? ", api.Method)
	}
	if api.Group != "" {
		db.Where("group = ? ", api.Group)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if orderString != "" {
			if desc {
				orderString = orderString + " desc"
			}
			err = db.Order(orderString, true).Find(&apiList).Error
		} else {
			err = db.Order("group", true).Find(&apiList).Error
		}
		return err, apiList, total
	}
}
