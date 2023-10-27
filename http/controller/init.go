package controller

import (
	"utopia-back/database/implement"
	v1 "utopia-back/http/controller/v1"
	v12 "utopia-back/service/implement/v1"
)

var (
	TestUserCtrl *v1.TestUserController
	UserCtrl     *v1.UserController
	FavoriteCtrl *v1.FavoriteController
	VideoCtrl    *v1.VideoController
	FollowCtrl   *v1.FollowController
)

func Init() {
	TestUserCtrl = v1.NewTestUserCtrl()
	UserCtrl = v1.NewUserController(&v12.UserService{Dal: &implement.UserDal{}})
	FavoriteCtrl = v1.NewFavoriteController(&v12.FavoriteService{Dal: &implement.FavoriteDal{}})
	VideoCtrl = v1.NewVideoController(&v12.VideoService{Dal: &implement.VideoImpl{}}) // todo
	FollowCtrl = v1.NewFollowController(&v12.FollowService{Dal: &implement.FollowDal{}})

}
