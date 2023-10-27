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

func (f *FollowDal) GetFansList(userId uint) (list []uint, err error) {
	res := database.DB.Where("follow_id = ?", userId).Select("user_id").Find(&list)
	return list, res.Error
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
