package router

import (
	"github.com/gin-gonic/gin"
	"utopia-back/database/implement"
	"utopia-back/http/controller"
	"utopia-back/http/middleware"
)

func Router(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimit)
	// 初始化Dal 保持单例
	centerDal := implement.NewCenterDal()

	v1ApiGroup := r.Group("/api/v1")
	{
		// 初始化控制器
		ctrlV1 := controller.NewCenterControllerV1(centerDal)

		// 测试模块
		testGroup := v1ApiGroup.Group("/test").Use(middleware.JwtMiddleware)
		{
			testGroup.POST("/testUser/add", ctrlV1.TestUserCtrl.Add)
			testGroup.GET("/testUser/select/:id", ctrlV1.TestUserCtrl.Select)
		}

		// 用户模块
		userGroup := v1ApiGroup.Group("/user")
		{
			// 登录注册
			userGroup.POST("/login", ctrlV1.UserCtrl.Login)
			userGroup.POST("/register", ctrlV1.UserCtrl.Register)

			// 修改昵称
			userGroup.POST("/nickname", middleware.JwtMiddleware, ctrlV1.UserCtrl.UpdateNickname)
			// 获取个人信息
			userGroup.GET("/info", middleware.JwtMiddleware, ctrlV1.UserCtrl.GetUserInfo)
		}

		// 交互模块
		interact := v1ApiGroup.Group("/interact").Use(middleware.JwtMiddleware)
		{
			interact.POST("/like", ctrlV1.LikeCtrl.Like)               // 点赞/取消点赞
			interact.POST("/follow", ctrlV1.FollowCtrl.Follow)         // 关注/取消关注
			interact.POST("/favorite", ctrlV1.FavoriteCtrl.Favorite)   // 收藏/取消收藏
			interact.GET("/follower/list", ctrlV1.FollowCtrl.FansList) // 获取粉丝列表
			interact.GET("/follow/list", ctrlV1.FollowCtrl.FollowList) // 获取关注列表
		}

		// 评论模块
		comment := v1ApiGroup.Group("/interact")
		{
			comment.POST("/comment", middleware.JwtMiddleware, ctrlV1.CommentCtrl.Comment)
			comment.GET("/comment/list", ctrlV1.CommentCtrl.CommentList)
		}

		// 存储模块
		storageGroup := v1ApiGroup.Group("/upload")
		{
			storageGroup.GET("/token", middleware.JwtMiddleware, ctrlV1.StorageCtrl.UploadToken)
			storageGroup.POST("/callback", ctrlV1.StorageCtrl.UploadCallback)
		}

		// 视频模块
		videoGroup := v1ApiGroup.Group("/video")
		{
			videoGroup.GET("/category", middleware.JwtWithoutAbortMiddleware, ctrlV1.VideoCtrl.GetCategoryVideos)
			videoGroup.GET("/popular", middleware.JwtWithoutAbortMiddleware, ctrlV1.VideoCtrl.GetPopularVideos)
			videoGroup.GET("/recommend", middleware.JwtWithoutAbortMiddleware, ctrlV1.VideoCtrl.GetRecommendVideos)
			videoGroup.GET("/favorite", middleware.JwtWithoutAbortMiddleware, ctrlV1.VideoCtrl.GetFavoriteVideos)
			videoGroup.GET("/upload", middleware.JwtWithoutAbortMiddleware, ctrlV1.VideoCtrl.GetUploadVideos)
			videoGroup.GET("/search", middleware.JwtWithoutAbortMiddleware, ctrlV1.VideoCtrl.SearchVideos)
			videoGroup.GET("/single", middleware.JwtWithoutAbortMiddleware, ctrlV1.VideoCtrl.GetSingleVideo)
		}
	}

	v2ApiGroup := r.Group("/api/v2")
	{
		// 初始化控制器
		ctrlV2 := controller.NewCenterControllerV2(centerDal)
		// 交互模块
		interact := v2ApiGroup.Group("/interact").Use(middleware.JwtMiddleware)
		{
			interact.POST("/follow", ctrlV2.FollowCtrl.Follow)         // 关注/取消关注
			interact.GET("/follower/list", ctrlV2.FollowCtrl.FansList) // 获取粉丝列表
			interact.GET("/follow/list", ctrlV2.FollowCtrl.FollowList) // 获取关注列表
			interact.POST("/like", ctrlV2.LikeCtrl.Like)               // 点赞/取消点赞
		}
	}

	v3ApiGroup := r.Group("/api/v3")
	{
		// 初始化控制器
		ctrlV3 := controller.NewCenterControllerV3(centerDal)
		// 交互模块
		interact := v3ApiGroup.Group("/interact").Use(middleware.JwtMiddleware)
		{
			interact.POST("/follow", ctrlV3.FollowCtrl.Follow)         // 关注/取消关注
			interact.GET("/follower/list", ctrlV3.FollowCtrl.FansList) // 获取粉丝列表
			interact.GET("/follow/list", ctrlV3.FollowCtrl.FollowList) // 获取关注列表
			interact.POST("/like", ctrlV3.LikeCtrl.Like)               // 点赞/取消点赞
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
