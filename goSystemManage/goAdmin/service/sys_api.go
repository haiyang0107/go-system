package service

import (
	"errors"
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
