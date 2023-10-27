package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
	"utopia-back/pkg/redis"
	abstract2 "utopia-back/service/abstract"
)

type FollowService struct {
	Dal abstract.FollowDal
}

func (f FollowService) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	return f.Dal.GetFollowList(userId)
}

// Follow 关注
func (f FollowService) Follow(userId uint, followId uint) (err error) {
	return f.Dal.Follow(userId, followId)
}

// UnFollow 取消关注
func (f FollowService) UnFollow(userId uint, followId uint) (err error) {
	return f.Dal.UnFollow(userId, followId)
}

// GetFansList 获取粉丝列表
func (f FollowService) GetFansList(userId uint) (list []model.UserInfo, err error) {
	return f.Dal.GetFansList(userId)
}

// IsFollow 是否关注
func (f FollowService) IsFollow(userId uint, followId uint) (isFollow bool, err error) {
	// 构造key
	key := "follow:" + "isFollow:" + string(userId) + ":" + string(followId)
	// 缓存
	res, err := redis.Cache(func() (interface{}, error) { return f.Dal.IsFollow(userId, followId) }, key, redis.TypeBool)
	// 返回结果
	if err != nil {
		return false, err
	}
	return res.(bool), nil
}

// GetFollowCount 获取关注数
func (f FollowService) GetFollowCount(userId uint) (count int64, err error) {
	// 构造key
	key := "follow:count:" + string(userId)
	// 缓存层取数据
	res, err := redis.Cache(func() (interface{}, error) { return f.Dal.GetFollowCount(userId) }, key, redis.TypeInt64)
	// 返回结果
	if err != nil {
		return 0, err
	}
	return res.(int64), nil

}

var _ abstract2.FollowService = (*FollowService)(nil)
