package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type FollowService struct {
	FollowDal abstract.FollowDal
	UserDal   abstract.UserDal
}

func (f FollowService) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	// 直接通过一条SQL查询出所有结果
	return f.FollowDal.GetFollowList(userId)
}

// Follow 关注
func (f FollowService) Follow(userId uint, followId uint) (err error) {
	// 判断followId是否存在
	_, err = f.UserDal.GetUserById(followId)
	if err != nil {
		return err
	}
	// 关注
	return f.FollowDal.Follow(userId, followId)
}

// UnFollow 取消关注
func (f FollowService) UnFollow(userId uint, followId uint) (err error) {
	// 判断followId是否存在
	_, err = f.UserDal.GetUserById(followId)
	if err != nil {
		return err
	}
	// 取消关注
	return f.FollowDal.UnFollow(userId, followId)
}

// GetFansList 获取粉丝列表
func (f FollowService) GetFansList(userId uint) (list []model.UserInfo, err error) {
	return f.FollowDal.GetFansList(userId)
}
