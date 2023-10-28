package router

import (
	"github.com/gin-gonic/gin"
	"utopia-back/http/controller"
	v12 "utopia-back/http/controller/v1"
	"utopia-back/http/middleware"
)

func Router(r *gin.Engine) *gin.Engine {

	v1ApiGroup := r.Group("/api/v1")

	// 初始化控制器
	ctrlV1 := controller.CenterController{
		TestUserCtrl: v12.NewTestUserCtrl(),
		UserCtrl:     v12.NewUserController(),
		FavoriteCtrl: v12.NewFavoriteController(),
		VideoCtrl:    v12.NewVideoController(),
		FollowCtrl:   v12.NewFollowController(),
	}

	// 测试模块
	testGroup := v1ApiGroup.Group("/test").Use(middleware.JwtMiddleware)
	{
		testGroup.POST("/testUser/add", ctrlV1.TestUserCtrl.Add)
		testGroup.GET("/testUser/select/:id", ctrlV1.TestUserCtrl.Select)
	}

	authGroup := v1ApiGroup.Group("/user")
	// 用户鉴权模块
	{
		authGroup.POST("/login", ctrlV1.UserCtrl.Login)
		authGroup.POST("/register", ctrlV1.UserCtrl.Register)
	}
	interact := v1ApiGroup.Group("/interact").Use(middleware.JwtMiddleware)
	// 交互模块
	{
		interact.POST("/favorite", ctrlV1.FavoriteCtrl.Favorite)   // 收藏/取消收藏
		interact.POST("/follow", ctrlV1.FollowCtrl.Follow)         // 关注/取消关注
		interact.GET("/follower/list", ctrlV1.FollowCtrl.FansList) // 获取粉丝列表
		interact.GET("/follow/list", ctrlV1.FollowCtrl.FollowList) // 获取关注列表
	}
	// 视频模块
	videoGroup := v1ApiGroup.Group("/video")
	{
		// 获取上传视频token
		videoGroup.GET("/upload/token", middleware.JwtMiddleware, ctrlV1.VideoCtrl.UploadVideoToken)
		videoGroup.POST("/upload/callback", ctrlV1.VideoCtrl.UploadVideoCallback)
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
