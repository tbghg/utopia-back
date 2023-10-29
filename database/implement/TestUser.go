package implement

import (
	"gorm.io/gorm"
	"utopia-back/model"
)

type TestUserDal struct{ Db *gorm.DB }

func (t *TestUserDal) Add(name string, age int) (id uint, err error) {
	newUser := model.TestUser{Name: name, Age: age}
	res := t.Db.Create(&newUser)
	if res.Error != nil {
		return 0, res.Error
	}
	return newUser.ID, nil
}

func (t *TestUserDal) Select(id uint) (user model.TestUser, err error) {
	res := t.Db.First(&user, id)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
