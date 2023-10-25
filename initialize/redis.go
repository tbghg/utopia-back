package initialize

import "utopia-back/pkg/redis"

func InitRedis() {
	// 初始化redis
	redis.Init()
}
