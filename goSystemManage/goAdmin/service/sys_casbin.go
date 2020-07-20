package service

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"goAdmin/global"
	"strings"
)

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

func ParamsMatch(name1 string, name2 string) bool {
	key := strings.Split(name1, "?")[0]
	return util.KeyMatch(key, name2)
}
