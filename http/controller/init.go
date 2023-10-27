package controller

import (
	"utopia-back/database/implement"
	v1 "utopia-back/http/controller/v1"
	v12 "utopia-back/service/implement/v1"
)

var (
	TestUserCtrl *v1.TestUserController
	UserCtrl     *v1.UserController
)

func Init() {
	TestUserCtrl = v1.NewTestUserCtrl()
	UserCtrl = v1.NewUserController(&v12.UserService{Dal: &implement.UserImpl{}})
}
