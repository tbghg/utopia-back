package main

import (
	"github.com/gin-gonic/gin"
	"utopia-back/config"
	"utopia-back/initialize"
)

func main() {
	// 初始化配置读取工具
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
	// 初始化redis
	initialize.InitRedis()
	// 初始化限流工具
	initialize.InitRateLimit()

	// 初始化配置
	r := gin.New()
	initialize.InitRoute(r)
	r.Run(config.V.GetString("server.port"))
}
