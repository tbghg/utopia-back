package implement

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type FollowDal struct {
	Db *gorm.DB
}

func (f *FollowDal) Follow(userId uint, followId uint) (err error) {
	res := f.Db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "follow_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"status": true}),
		}).Create(&model.Follow{
		UserID:   userId,
		FollowID: followId,
		Status:   true,
	})
	return res.Error
}

func (f *FollowDal) UnFollow(userId uint, followId uint) (err error) {
	res := f.Db.Model(&model.Follow{}).Where("user_id = ? AND follow_id = ?", userId, followId).Update("status", false)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return err
}

func (f *FollowDal) GetFansList(userId uint) (list []model.UserInfo, err error) {
	var users []model.UserInfo
	//联表查询
	res := f.Db.Model(model.Follow{}).
		Select("users.*, COUNT(DISTINCT following.id) AS follow_count, COUNT(DISTINCT followers.id) AS fans_count").
		Joins("LEFT JOIN users ON users.id = follows.user_id").
		Joins("LEFT JOIN follows AS following ON following.user_id = users.id").
		Joins("LEFT JOIN follows AS followers ON followers.follow_id = users.id").
		Where("follows.follow_id IN (?)", userId).
		Group("follows.user_id").
		Find(&users)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return users, err
}

// GetFollowList 获取关注列表
func (f *FollowDal) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	var users []model.UserInfo
	//联表查询
	res := f.Db.Model(model.Follow{}).
		Select("users.*, COUNT(DISTINCT following.id) AS follow_count, COUNT(DISTINCT followers.id) AS fans_count").
		Joins("LEFT JOIN users ON users.id = follows.follow_id").
		Joins("LEFT JOIN follows AS following ON following.user_id = users.id").
		Joins("LEFT JOIN follows AS followers ON followers.follow_id = users.id").
		Where("follows.user_id IN (?)", userId).
		Group("follows.follow_id").
		Find(&users)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return users, err
}

func (f *FollowDal) IsFollow(userId uint, followId uint) (isFollow bool, err error) {
	var follow model.Follow
	res := f.Db.Where("user_id = ? AND follow_id = ?", userId, followId).First(&follow)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return follow.Status, err
}

func (f *FollowDal) GetFollowCount(userId uint) (count int64, err error) {
	var follow model.Follow
	res := f.Db.Where("user_id = ?", userId).Find(&follow).Count(&count)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return count, err
}

func (f *FollowDal) GetFansCount(userId uint) (count int64, err error) {
	var follow model.Follow
	res := f.Db.Where("follow_id = ?", userId).Find(&follow).Count(&count)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return count, err
}

// GetFansIdList 获取粉丝id列表
func (f *FollowDal) GetFansIdList(userId uint) (fansIdList []uint, err error) {
	res := f.Db.Model(&model.Follow{}).Select("user_id").Where("follow_id = ?", userId).Find(&fansIdList)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return fansIdList, err
}

// GetFollowIdList 获取自己关注的用户id列表
func (f *FollowDal) GetFollowIdList(userId uint) (followIdList []uint, err error) {
	res := f.Db.Model(&model.Follow{}).Select("follow_id").Where("user_id = ?", userId).Find(&followIdList)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return followIdList, err
}

var _ abstract.FollowDal = (*FollowDal)(nil)
