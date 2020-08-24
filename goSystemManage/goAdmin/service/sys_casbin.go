package service

//noinspection GoUnresolvedReference
import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"goAdmin/global"
	"goAdmin/model"
	"strings"
)

//noinspection GoUnresolvedReference
func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(global.GLOBAL_DB)
	e := casbin.NewEnforcer(global.GLOBAL_CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	e.LoadPolicy()
	return e
}

func ParamsMatchFunc(args ...interface{}) (i interface{}, err error) {
	firstName := args[0].(string)
	secondName := args[1].(string)
	return ParamsMatch(firstName, secondName), nil
}

//@title  ParamsMatch
//description  customized rule ,自定义规则函数

//noinspection GoUnresolvedReference
func ParamsMatch(name1 string, name2 string) bool {
	key := strings.Split(name1, "?")[0]
	return util.KeyMatch(key, name2)
}

//更新权限组内的请求api信息
func UpdateCasbinApi(oldPath string, oldMethod string, newPath string, newMethod string) (err error) {
	var cbm []model.SysCasbin
	return global.GLOBAL_DB.Table("sys_casbins").Where("path = ? AND method = ? ", oldPath, oldMethod).Find(&cbm).Update(map[string]string{
		"path":   newPath,
		"method": newMethod,
	}).Error
}

//清楚匹配的权限
func clearCabinApi(v int, p ...string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p...)
}
