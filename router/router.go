package router

import (
	"github.com/gin-gonic/gin"
	v1 "utopia-back/http/controller/v1"
)

func Router(r *gin.Engine) *gin.Engine {

	v1ApiGroup := r.Group("/api/v1")
	testUserCtrl := v1.NewTestUserCtrl()
	{
		v1ApiGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1ApiGroup.POST("/testUser/add", testUserCtrl.Add)
		v1ApiGroup.GET("/testUser/select/:id", testUserCtrl.Select)
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
