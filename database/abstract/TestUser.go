package abstract

import "utopia-back/model"

type TestUserDal interface {
	Add(name string, age int) (id uint, err error)
	Select(id uint) (user model.TestUser, err error)
}
