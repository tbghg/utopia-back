package abstract

import "utopia-back/model"

type FollowDal interface {
	// Follow 关注
	Follow(userId uint, followId uint) (err error)
	// UnFollow 取消关注
	UnFollow(userId uint, followId uint) (err error)
	// IsFollow 是否关注
	IsFollow(userId uint, followId uint) (isFollow bool, err error)
	// GetFollowCount 获取关注数
	GetFollowCount(userId uint) (count int64, err error)
	// GetFansCount 获取粉丝数
	GetFansCount(userId uint) (count int64, err error)
	// GetFansIdList 获取粉丝id列表
	GetFansIdList(userId uint) (fansIdList []uint, err error)
	// GetFollowIdList 获取关注id列表
	GetFollowIdList(userId uint) (followIdList []uint, err error)
	// GetFollowList 获取关注列表
	GetFollowList(userId uint) (list []model.UserInfo, err error)
	// GetFansList 获取粉丝列表
	GetFansList(userId uint) (list []model.UserInfo, err error)
}
