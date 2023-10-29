package router

import (
	"github.com/gin-gonic/gin"
	"utopia-back/database/implement"
	"utopia-back/http/controller"
	v12 "utopia-back/http/controller/v1"
	v2 "utopia-back/http/controller/v2"
	v3 "utopia-back/http/controller/v3"
	"utopia-back/http/middleware"
	v1 "utopia-back/service/implement/v1"
	v22 "utopia-back/service/implement/v2"
	v32 "utopia-back/service/implement/v3"
)

func Router(r *gin.Engine) *gin.Engine {
	r.Use(middleware.RateLimit)
	// 初始化Dal 保持单例
	testUserDal := &implement.TestUserDal{}
	userDal := &implement.UserDal{}
	favoriteDal := &implement.FavoriteDal{}
	followDal := &implement.FollowDal{}
	likeDal := &implement.LikeDal{}
	videoDal := &implement.VideoDal{}

	v1ApiGroup := r.Group("/api/v1")
	{
		// 初始化控制器
		ctrlV1 := controller.CenterControllerV1{
			// TestUserController
			TestUserCtrl: &v12.TestUserController{
				TestUserService: &v1.TestUserService{
					TestUserDal: testUserDal,
				},
			},
			// UserController
			UserCtrl: &v12.UserController{
				UserService: &v1.UserService{
					UserDal: userDal,
				},
			},
			// FavoriteController
			FavoriteCtrl: &v12.FavoriteController{
				FavoriteService: &v1.FavoriteService{
					FavoriteDal: favoriteDal,
					VideoDal:    videoDal,
				},
			},
			// StorageController
			StorageCtrl: &v12.StorageController{
				StorageService: &v1.StorageService{
					VideoDal: videoDal,
					UserDal:  userDal,
				},
			},
			// FollowController
			FollowCtrl: &v12.FollowController{
				FollowService: &v1.FollowService{
					FollowDal: followDal,
					UserDal:   userDal,
				},
			},
			// LikeController
			LikeCtrl: &v12.LikeController{
				LikeService: &v1.LikeService{
					LikeDal:  likeDal,
					VideoDal: videoDal,
				},
			},
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
			interact.POST("/like", ctrlV1.LikeCtrl.Like)               // 点赞/取消点赞
			interact.POST("/follow", ctrlV1.FollowCtrl.Follow)         // 关注/取消关注
			interact.POST("/favorite", ctrlV1.FavoriteCtrl.Favorite)   // 收藏/取消收藏
			interact.GET("/follower/list", ctrlV1.FollowCtrl.FansList) // 获取粉丝列表
			interact.GET("/follow/list", ctrlV1.FollowCtrl.FollowList) // 获取关注列表
		}
		// 存储模块
		storageGroup := v1ApiGroup.Group("/upload")
		{
			storageGroup.GET("/token", middleware.JwtMiddleware, ctrlV1.StorageCtrl.UploadToken)
			storageGroup.POST("/callback", ctrlV1.StorageCtrl.UploadCallback)
		}
	}

	v2ApiGroup := r.Group("/api/v2")
	{
		// 初始化控制器
		ctrlV2 := controller.CenterControllerV2{
			FollowCtrl: &v2.FollowController{
				FollowService: v22.FollowService{
					FollowDal: &implement.FollowDal{},
				},
			},
		}
		// 交互模块
		interact := v2ApiGroup.Group("/interact").Use(middleware.JwtMiddleware)
		{
			interact.POST("/follow", ctrlV2.FollowCtrl.Follow)         // 关注/取消关注
			interact.GET("/follower/list", ctrlV2.FollowCtrl.FansList) // 获取粉丝列表
			interact.GET("/follow/list", ctrlV2.FollowCtrl.FollowList) // 获取关注列表
		}
	}

	v3ApiGroup := r.Group("/api/v3")
	{
		// 初始化控制器
		ctrlV3 := controller.CenterControllerV3{
			FollowCtrl: &v3.FollowController{
				FollowService: &v32.FollowService{
					FollowDal: &implement.FollowDal{},
					UserDal:   &implement.UserDal{},
				},
			},
		}
		// 交互模块
		interact := v3ApiGroup.Group("/interact").Use(middleware.JwtMiddleware)
		{
			interact.POST("/follow", ctrlV3.FollowCtrl.Follow)         // 关注/取消关注
			interact.GET("/follower/list", ctrlV3.FollowCtrl.FansList) // 获取粉丝列表
			interact.GET("/follow/list", ctrlV3.FollowCtrl.FollowList) // 获取关注列表
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
