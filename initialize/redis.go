package initialize

import "utopia-back/cache"

func InitRedis() {
	// 初始化redis
	cache.Init()
}
