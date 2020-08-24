package main

import (
	"fmt"
	"goAdmin/config"
	"goAdmin/global"
	"goAdmin/initialization"
)

func main() {
	config.Init()
	switch global.GLOBAL_CONFIG.System.DbType {
	case "mysql":
		fmt.Println("初始化MySQL")
		initialization.Mysql()
	default:
		initialization.Mysql()
		fmt.Println("默认初始化MySQL")
	}
	//初始化表结构
	initialization.InitDbTables()
	fmt.Printf("this is a new start about go!")
	//开启服务
	config.RunServer()
}
