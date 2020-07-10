package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"goAdmin/base"
)

//定义全局变量
var (
	GLOBAL_DB     *gorm.DB
	GLOBAL_REDIS  *redis.Client //后期加入redis服务
	GLOBAL_CONFIG base.Server
	GLOBAL_VP     *viper.Viper
	GLOBAL_LOG    *oplogging.Logger
)
