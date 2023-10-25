package controller

import v1 "utopia-back/http/controller/v1"

var (
	TestUserCtrl *v1.TestUserController
	UserCtrl     *v1.UserController
)

func Init() {
	TestUserCtrl = v1.NewTestUserCtrl()
	UserCtrl = v1.NewUserController()
}
