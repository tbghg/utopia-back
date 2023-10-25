package abstract

import "utopia-back/model"

type UserDal interface {
	GetUserByUsername(username string) (user model.User, err error)
	CreateUser(user *model.User) (id uint, err error)
}
