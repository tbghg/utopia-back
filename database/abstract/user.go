package abstract

import "utopia-back/model"

type UserDal interface {
	GetUserByUsername(username string) (user model.User, exist bool, err error)
	CreateUser(user *model.User) (id uint, err error)
	GetUserById(id uint) (user model.User, err error)
	GetUserInfoById(id uint) (userInfo model.UserInfo, err error)
	UpdateAvatar(id uint, avatarUrl string) (err error)
	UpdateNickname(id uint, nickname string) (err error)
}
