package implement

import (
	"errors"
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type UserDal struct {
	Db *gorm.DB
}

func (u *UserDal) GetUserInfoById(id uint) (userInfo model.UserInfo, err error) {
	userInfo = model.UserInfo{}
	userInfo.ID = id

	// 获取用户基本信息
	var user model.User
	user, err = u.GetUserById(id)
	if err != nil {
		return userInfo, err
	}
	userInfo.Username = user.Username
	userInfo.Nickname = user.Nickname
	userInfo.Avatar = user.Avatar

	// 获取用户的关注数
	var followCount int64
	res := u.Db.Model(&model.Follow{}).Where("user_id = ?", id).Count(&followCount)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	if err != nil {
		return userInfo, err
	}
	userInfo.FollowCount = followCount

	// 获取用户的粉丝数
	var fansCount int64
	res = u.Db.Model(&model.Follow{}).Where("follow_id = ?", id).Count(&fansCount)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	if err != nil {
		return userInfo, err
	}
	userInfo.FansCount = fansCount

	// 获取用户的视频作品数
	var videoCount int64
	res = u.Db.Model(&model.Video{}).Where("author_id = ?", id).Count(&videoCount)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	if err != nil {
		return userInfo, err
	}
	userInfo.VideoCount = videoCount

	return userInfo, nil

}

func (u *UserDal) GetUserByUsername(username string) (user model.User, exist bool, err error) {
	res := u.Db.First(&user, "username = ?", username)
	if res.Error == nil {
		exist = true
	}
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return
}

func (u *UserDal) CreateUser(user *model.User) (id uint, err error) {
	res := u.Db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

func (u *UserDal) GetUserById(id uint) (user model.User, err error) {
	res := u.Db.First(&user, "id = ?", id)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return user, err
}

func (u *UserDal) UpdateAvatar(id uint, avatarUrl string) (err error) {
	res := u.Db.Model(&model.User{ID: id}).Update("avatar", avatarUrl)
	if res.Error != nil {
		return res.Error
	}
	return
}

func (u *UserDal) UpdateNickname(id uint, nickname string) (err error) {
	res := u.Db.Model(&model.User{ID: id}).Update("nickname", nickname)
	if res.Error != nil {
		return res.Error
	}
	return
}

var _ abstract.UserDal = (*UserDal)(nil)
