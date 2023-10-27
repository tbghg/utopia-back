package abstract

import "utopia-back/model"

type FollowService interface {
	// Follow 关注
	Follow(userId uint, followId uint) (err error)
	// UnFollow 取消关注
	UnFollow(userId uint, followId uint) (err error)
	// GetFansList 获取粉丝列表
	GetFansList(userId uint) (list []model.UserInfo, err error)
	// IsFollow 是否关注
	IsFollow(userId uint, followId uint) (isFollow bool, err error)
	// GetFollowCount 获取关注数
	GetFollowCount(userId uint) (count int64, err error)
	// GetFollowList 获取关注列表
	GetFollowList(userId uint) (list []model.UserInfo, err error)
}
