package service

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type TestUserService struct {
	Dal abstract.TestUserDal
}

func (t *TestUserService) Add(name string, age int) (id uint, err error) {
	return t.Dal.Add(name, age)
}

func (t *TestUserService) Select(id uint) (user model.TestUser, err error) {
	return t.Dal.Select(id)
}
