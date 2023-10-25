package router

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {

	v1ApiGroup := r.Group("/api/v1")

	{
		v1ApiGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	return r
}

func Handle404Route(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    "404",
			"message": "路由不存在,请检查请求方法和请求路径",
		})
	})
}
