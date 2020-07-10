package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goAdmin/global"
)

//读取配置文件信息，并启动项目的基类
const defaultConfigFile = "config.yml"

//初始化读取配置文件
func Init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GLOBAL_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GLOBAL_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GLOBAL_VP = v
}

//启动项目server
func RunServer() {
	fmt.Println("a big go system is run with server!")
}
