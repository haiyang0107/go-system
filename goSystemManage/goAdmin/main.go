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

	fmt.Printf("this is a new start about go!")
	config.RunServer()
}
