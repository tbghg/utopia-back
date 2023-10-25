package implement

import (
	"utopia-back/database"
	"utopia-back/model"
)

type UserImpl struct {
}

func (u *UserImpl) GetUserByUsername(username string) (user model.User, err error) {
	res := database.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (u *UserImpl) CreateUser(user *model.User) (id uint, err error) {
	res := database.DB.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}
