package router

import (
	"github.com/gin-gonic/gin"
	"utopia-back/http/controller"
)

func Router(r *gin.Engine) *gin.Engine {

	v1ApiGroup := r.Group("/api/v1")
	authGroup := v1ApiGroup.Group("/auth")

	{
		v1ApiGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1ApiGroup.POST("/testUser/add", controller.TestUserCtrl.Add)
		v1ApiGroup.GET("/testUser/select/:id", controller.TestUserCtrl.Select)
		//鉴权模块
		{
			authGroup.POST("/login", controller.UserCtrl.Login)
			authGroup.POST("/register", controller.UserCtrl.Register)
		}
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
