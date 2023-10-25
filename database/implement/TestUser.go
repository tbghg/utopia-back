package implement

import (
	"utopia-back/database"
	"utopia-back/model"
)

type TestUserImpl struct{}

func (t *TestUserImpl) Add(name string, age int) (id uint, err error) {
	newUser := model.TestUser{Name: name, Age: age}
	res := database.DB.Create(&newUser)
	if res.Error != nil {
		return 0, res.Error
	}
	return newUser.ID, nil
}

func (t *TestUserImpl) Select(id uint) (user model.TestUser, err error) {
	res := database.DB.First(&user, id)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
