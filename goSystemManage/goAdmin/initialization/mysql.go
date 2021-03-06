package initialization

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goAdmin/global"
	"goAdmin/model"
	"os"
)

//初始化MySQL数据库并配置全局变量
func Mysql() {
	admin := global.GLOBAL_CONFIG.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		global.GLOBAL_LOG.Error("mysql 启动异常", err)
		os.Exit(0)
	} else {
		global.GLOBAL_DB = db
		global.GLOBAL_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		global.GLOBAL_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.GLOBAL_DB.LogMode(admin.LogMode)
		global.GLOBAL_LOG.Error("mysql 启动成功")
	}
}

//初始化表结构，注册数据库专用
func InitDbTables() {
	db := global.GLOBAL_DB
	db.AutoMigrate(
		model.SysUser{},
		model.SysApi{},
		model.SysMenu{},
		model.SysRole{},
		model.SysJwtBlack{},
		model.SysCasbin{},
	)
	global.GLOBAL_LOG.Debug("register table   success")
}
