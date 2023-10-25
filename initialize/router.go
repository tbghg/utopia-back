package initialize

import (
	"github.com/gin-gonic/gin"
	"utopia-back/http/middleware"
	"utopia-back/router"
)

// InitRoute 初始化路由
func InitRoute(r *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleware(r)
	//注册API路由
	router.Router(r)
	//404处理
	router.Handle404Route(r)
}

// registerGlobalMiddleware 注册全局中间件
func registerGlobalMiddleware(r *gin.Engine) {
	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
	)
}
