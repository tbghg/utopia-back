package initialize

import (
	"github.com/gin-gonic/gin"
	"utopia-back/router"
)

// InitRoute 初始化路由
func InitRoute(r *gin.Engine) {
	//注册API路由
	router.Router(r)
	//404处理
	router.Handle404Route(r)
}
