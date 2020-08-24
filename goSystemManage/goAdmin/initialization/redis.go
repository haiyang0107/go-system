package initialization

import (
	"github.com/go-redis/redis"
	"goAdmin/global"
)

func Redis() {
	//获取redis配置信息
	redisconfig := global.GLOBAL_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisconfig.Addr,
		Password: redisconfig.Password,
		DB:       redisconfig.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GLOBAL_LOG.Error(err)
	} else {
		global.GLOBAL_LOG.Info("redis connect ping response", pong)
		global.GLOBAL_REDIS = client
	}
}
