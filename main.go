package main

import (
	"github.com/gin-gonic/gin"
	"utopia-back/initialize"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
	// 初始化redis
	initialize.InitRedis()

	// 初始化配置
	r := gin.New()
	initialize.InitRoute(r)
	r.Run()
}
