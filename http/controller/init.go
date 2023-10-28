package controller

import (
	v1 "utopia-back/http/controller/v1"
)

type CenterController struct {
	TestUserCtrl *v1.TestUserController
	UserCtrl     *v1.UserController
	FavoriteCtrl *v1.FavoriteController
	StorageCtrl  *v1.StorageController
	FollowCtrl   *v1.FollowController
}
