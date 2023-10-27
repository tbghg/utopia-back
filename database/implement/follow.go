package implement

import (
	"utopia-back/database"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type FollowDal struct{}

var _ abstract.FollowDal = (*FollowDal)(nil)

func (f *FollowDal) Follow(userId uint, followId uint) (err error) {
	res := database.DB.Create(&model.Follow{
		UserID: userId,
		FunID:  followId,
	})
	return res.Error
}

func (f *FollowDal) UnFollow(userId uint, followId uint) (err error) {
	res := database.DB.Where("user_id = ? AND follow_id = ?", userId, followId).Delete(&model.Follow{})
	return res.Error
}

func (f *FollowDal) GetFansList(userId uint) (list []model.UserInfo, err error) {
	var users []model.UserInfo
	//联表查询
	res := database.DB.Model(model.Follow{}).
		Select("users.*, COUNT(DISTINCT following.id) AS follow_count, COUNT(DISTINCT followers.id) AS fans_count").
		Joins("LEFT JOIN users ON users.id = follows.user_id").
		Joins("LEFT JOIN follows AS following ON following.user_id = users.id").
		Joins("LEFT JOIN follows AS followers ON followers.fun_id = users.id").
		Where("follows.fun_id IN (?)", userId).
		Group("follows.user_id").
		Find(&users)

	return users, res.Error
}

// GetFollowList 获取关注列表
func (f *FollowDal) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	var users []model.UserInfo
	//联表查询
	res := database.DB.Model(model.Follow{}).
		Select("users.*, COUNT(DISTINCT following.id) AS follow_count, COUNT(DISTINCT followers.id) AS fans_count").
		Joins("LEFT JOIN users ON users.id = follows.fun_id").
		Joins("LEFT JOIN follows AS following ON following.user_id = users.id").
		Joins("LEFT JOIN follows AS followers ON followers.fun_id = users.id").
		Where("follows.user_id IN (?)", userId).
		Group("follows.fun_id").
		Find(&users)

	return users, res.Error
}

func (f *FollowDal) IsFollow(userId uint, followId uint) (isFollow bool, err error) {
	var follow model.Follow
	res := database.DB.Where("user_id = ? AND follow_id = ?", userId, followId).First(&follow)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (f *FollowDal) GetFollowCount(userId uint) (count int64, err error) {
	var follow model.Follow
	res := database.DB.Where("user_id = ?", userId).Find(&follow).Count(&count)
	return count, res.Error
}

func (f *FollowDal) GetFansCount(userId uint) (count int64, err error) {
	var follow model.Follow
	res := database.DB.Where("follow_id = ?", userId).Find(&follow).Count(&count)
	return count, res.Error
}
