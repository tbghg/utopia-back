package controller

import (
	v1 "utopia-back/http/controller/v1"
	v2 "utopia-back/http/controller/v2"
)

type CenterControllerV1 struct {
	TestUserCtrl *v1.TestUserController
	UserCtrl     *v1.UserController
	FavoriteCtrl *v1.FavoriteController
	StorageCtrl  *v1.StorageController
	FollowCtrl   *v1.FollowController
}

type CenterControllerV2 struct {
	FollowCtrl *v2.FollowController
}
