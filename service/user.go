package service

import (
	"errors"
	"utopia-back/database/abstract"
	"utopia-back/model"
	utils "utopia-back/pkg/util"
)

type UserService struct {
	Dal abstract.UserDal
}

var (
	ErrorUsernameEmpty = errors.New("用户名和密码不能为空")
	ErrorUserExists    = errors.New("用户已存在")
	ErrorUserNotExists = errors.New("用户不存在")
	ErrorPasswordWrong = errors.New("密码错误")
)

func (u *UserService) Login(username string, password string) (token string, id uint, err error) {
	// 检查用户名和密码是否为空
	if username == "" || password == "" {
		return "", 0, ErrorUsernameEmpty
	}
	// 检查用户是否存在
	user, err := u.Dal.GetUserByUsername(username)
	if err != nil {
		return "", 0, ErrorUserNotExists
	}
	// 检查密码是否正确
	if utils.ValidMd5EncodeWithSalt(password, user.Salt, user.Password) {
		// todo 生成token
		return "token", user.ID, nil
	}
	return "", 0, ErrorPasswordWrong
}

func (u *UserService) Register(username string, password string) (token string, id uint, err error) {
	// 检查用户名和密码是否为空
	if username == "" || password == "" {
		return "", 0, ErrorUsernameEmpty
	}
	// 检查用户是否存在
	_, err = u.Dal.GetUserByUsername(username)
	if err == nil {
		return "", 0, ErrorUserExists
	}

	// 密码加密+盐
	salt := utils.RandomSalt()
	password = utils.Md5EncodeWithSalt(password, salt)
	user := model.User{Username: username, Password: password, Salt: salt}
	// 创建用户
	id, err = u.Dal.CreateUser(&user)
	if err != nil {
		return "", 0, err
	}
	// todo 生成token
	return "token", id, err
}
