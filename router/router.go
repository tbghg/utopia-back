package router

import (
	"github.com/gin-gonic/gin"
	"utopia-back/http/controller"
	"utopia-back/http/middleware"
)

func Router(r *gin.Engine) *gin.Engine {

	v1ApiGroup := r.Group("/api/v1")
	authGroup := v1ApiGroup.Group("/auth")
	testGroup := v1ApiGroup.Group("/test").Use(middleware.JwtMiddleware)

	{
		v1ApiGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		//鉴权模块
		{
			authGroup.POST("/login", controller.UserCtrl.Login)
			authGroup.POST("/register", controller.UserCtrl.Register)
		}
		{
			testGroup.POST("/testUser/add", controller.TestUserCtrl.Add)
			testGroup.GET("/testUser/select/:id", controller.TestUserCtrl.Select)
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
