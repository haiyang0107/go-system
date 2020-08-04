package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goAdmin/global"
	"goAdmin/initialization"
	"net/http"
	"time"
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
	//添加服务注册端口
	Router := initialization.Routers()
	Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.GLOBAL_CONFIG.System.Addr)
	h := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	global.GLOBAL_LOG.Info("server port on :", h.Addr)
	fmt.Printf("欢迎使用 goSystemMange 	发起者 ： 逆行者工作室	微信 ： HyOcean07")
	h.ListenAndServe()
}
