package router

import (
	"github.com/gin-gonic/gin"
	"utopia-back/http/controller"
	"utopia-back/http/middleware"
)

func Router(r *gin.Engine) *gin.Engine {

	v1ApiGroup := r.Group("/api/v1")
	authGroup := v1ApiGroup.Group("/user")
	testGroup := v1ApiGroup.Group("/test").Use(middleware.JwtMiddleware)
	interact := v1ApiGroup.Group("/interact").Use(middleware.JwtMiddleware)

	{

		testGroup.POST("/testUser/add", controller.TestUserCtrl.Add)
		testGroup.GET("/testUser/select/:id", controller.TestUserCtrl.Select)
		// 用户鉴权模块
		{
			authGroup.POST("/login", controller.UserCtrl.Login)
			authGroup.POST("/register", controller.UserCtrl.Register)
		}
		// 交互模块
		{
			interact.POST("/favorite", controller.FavoriteCtrl.Favorite)
			interact.POST("/follow", controller.FollowCtrl.Follow)
			interact.GET("/follower/list", controller.FollowCtrl.FansList)
		}

	}

	videoGroup := v1ApiGroup.Group("/video")
	{
		// 获取上传视频token
		videoGroup.GET("/upload/token", controller.VideoCtrl.UploadVideoToken).Use(middleware.JwtMiddleware)
		videoGroup.GET("/upload/callback", controller.VideoCtrl.UploadVideoCallback)

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
