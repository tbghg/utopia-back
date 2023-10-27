package implement

import (
	"utopia-back/database"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type UserDal struct {
}

func (u *UserDal) GetUserInfoById(id uint) (userInfo model.UserInfo, err error) {
	userInfo = model.UserInfo{}
	userInfo.ID = id
	// 获取用户基本信息
	user, err := u.GetUserById(id)
	if err != nil {
		return userInfo, err
	}
	userInfo.Username = user.Username
	userInfo.Nickname = user.Nickname
	userInfo.Avatar = user.Avatar
	// 获取用户的关注数
	var followCount int64
	res := database.DB.Model(&model.Follow{}).Where("user_id = ?", id).Count(&followCount)
	if res.Error != nil {
		return userInfo, res.Error
	}
	userInfo.FollowCount = followCount
	// 获取用户的粉丝数
	var fansCount int64
	res = database.DB.Model(&model.Follow{}).Where("follow_id = ?", id).Count(&fansCount)
	if res.Error != nil {
		return userInfo, res.Error
	}
	userInfo.FansCount = fansCount
	// 获取用户的视频作品数
	var videoCount int64
	res = database.DB.Model(&model.Video{}).Where("user_id = ?", id).Count(&videoCount)
	if res.Error != nil {
		return userInfo, res.Error
	}
	userInfo.VideoCount = videoCount

	return userInfo, nil

}

func (u *UserDal) GetUserByUsername(username string) (user model.User, err error) {
	res := database.DB.First(&user, "username = ?", username)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (u *UserDal) CreateUser(user *model.User) (id uint, err error) {
	res := database.DB.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

func (u *UserDal) GetUserById(id uint) (user model.User, err error) {
	res := database.DB.First(&user, "id = ?", id)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

var _ abstract.UserDal = (*UserDal)(nil)
