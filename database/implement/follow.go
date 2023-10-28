package implement

import (
	"gorm.io/gorm/clause"
	"utopia-back/database"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type FollowDal struct{}

func (f *FollowDal) Follow(userId uint, followId uint) (err error) {
	res := database.DB.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "follow_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"status": true}),
		}).Create(&model.Follow{
		UserID: userId,
		FunID:  followId,
		Status: true,
	})
	return res.Error
}

func (f *FollowDal) UnFollow(userId uint, followId uint) (err error) {
	res := database.DB.Model(&model.Follow{}).Where("user_id = ? AND fun_id = ?", userId, followId).Update("status", false)
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
	res := database.DB.Where("fun_id = ?", userId).Find(&follow).Count(&count)
	return count, res.Error
}

// GetFansIdList 获取粉丝id列表
func (f *FollowDal) GetFansIdList(userId uint) (fansIdList []uint, err error) {
	res := database.DB.Model(&model.Follow{}).Select("user_id").Where("fun_id = ?", userId).Find(&fansIdList)
	if res.Error != nil {
		return nil, res.Error
	}
	return fansIdList, nil
}

// GetFollowIdList 获取自己关注的用户id列表
func (f *FollowDal) GetFollowIdList(userId uint) (followIdList []uint, err error) {
	res := database.DB.Model(&model.Follow{}).Select("fun_id").Where("user_id = ?", userId).Find(&followIdList)
	if res.Error != nil {
		return nil, res.Error
	}
	return followIdList, nil
}

var _ abstract.FollowDal = (*FollowDal)(nil)
