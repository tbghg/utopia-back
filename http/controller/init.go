package controller

import (
	"utopia-back/database/implement"
	v1 "utopia-back/http/controller/v1"
	v2 "utopia-back/http/controller/v2"
	v3 "utopia-back/http/controller/v3"
	v12 "utopia-back/service/implement/v1"
	v22 "utopia-back/service/implement/v2"
	v32 "utopia-back/service/implement/v3"
)

type CenterControllerV1 struct {
	TestUserCtrl *v1.TestUserController
	UserCtrl     *v1.UserController
	FavoriteCtrl *v1.FavoriteController
	StorageCtrl  *v1.StorageController
	FollowCtrl   *v1.FollowController
	LikeCtrl     *v1.LikeController
	VideoCtrl    *v1.VideoController
}

type CenterControllerV2 struct {
	FollowCtrl *v2.FollowController
	LikeCtrl   *v2.LikeController
}

type CenterControllerV3 struct {
	FollowCtrl *v3.FollowController
}

func NewCenterControllerV1(dal *implement.CenterDal) *CenterControllerV1 {
	testUserCtrl := &v1.TestUserController{
		TestUserService: &v12.TestUserService{TestUserDal: dal.TestUserDal},
	}
	userCtrl := &v1.UserController{
		UserService: &v12.UserService{UserDal: dal.UserDal},
	}
	favoriteCtrl := &v1.FavoriteController{
		FavoriteService: &v12.FavoriteService{FavoriteDal: dal.FavoriteDal},
	}
	storageCtrl := &v1.StorageController{
		StorageService: &v12.StorageService{UserDal: dal.UserDal, VideoDal: dal.VideoDal},
	}
	followCtrl := &v1.FollowController{
		FollowService: &v12.FollowService{FollowDal: dal.FollowDal, UserDal: dal.UserDal},
	}
	likeCtrl := &v1.LikeController{
		LikeService: &v12.LikeService{LikeDal: dal.LikeDal, VideoDal: dal.VideoDal},
	}
	videoCtrl := &v1.VideoController{
		VideoService: &v12.VideoService{
			VideoDal:    dal.VideoDal,
			UserDal:     dal.UserDal,
			FollowDal:   dal.FollowDal,
			LikeDal:     dal.LikeDal,
			FavoriteDal: dal.FavoriteDal,
		},
	}

	return &CenterControllerV1{
		TestUserCtrl: testUserCtrl,
		UserCtrl:     userCtrl,
		FavoriteCtrl: favoriteCtrl,
		StorageCtrl:  storageCtrl,
		FollowCtrl:   followCtrl,
		LikeCtrl:     likeCtrl,
		VideoCtrl:    videoCtrl,
	}
}

func NewCenterControllerV2(dal *implement.CenterDal) *CenterControllerV2 {
	followCtrl := &v2.FollowController{
		FollowService: &v22.FollowService{FollowDal: dal.FollowDal, UserDal: dal.UserDal},
	}

	likeCtrl := &v2.LikeController{
		LikeService: &v22.LikeService{LikeDal: dal.LikeDal},
	}

	return &CenterControllerV2{
		FollowCtrl: followCtrl,
		LikeCtrl:   likeCtrl,
	}
}

func NewCenterControllerV3(dal *implement.CenterDal) *CenterControllerV3 {
	followCtrl := &v3.FollowController{
		FollowService: &v32.FollowService{FollowDal: dal.FollowDal, UserDal: dal.UserDal},
	}

	return &CenterControllerV3{
		FollowCtrl: followCtrl,
	}
}
