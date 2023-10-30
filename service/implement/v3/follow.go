package v3

import (
	"errors"
	"utopia-back/database/abstract"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type FollowService struct {
	FollowDal abstract.FollowDal
	UserDal   abstract.UserDal
}

var ErrUserNotExist = errors.New("关注的用户不存在")

// Follow 关注
func (f FollowService) Follow(userId uint, followId uint) (err error) {
	return f.FollowDal.Follow(userId, followId)
}

// UnFollow 取消关注
func (f FollowService) UnFollow(userId uint, followId uint) (err error) {
	return f.FollowDal.UnFollow(userId, followId)
}

// GetFansList 获取粉丝列表
func (f FollowService) GetFansList(userId uint) (list []model.UserInfo, err error) {
	//先获取粉丝id列表
	fansIdList, err := f.FollowDal.GetFansIdList(userId)
	if err != nil {
		return nil, err
	}
	// 再根据粉丝id列表获取粉丝信息列表
	return f.getUserInfo(fansIdList)
}

// GetFollowList 获取关注列表
func (f FollowService) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	//先获取自己关注的用户id列表
	followIdList, err := f.FollowDal.GetFollowIdList(userId)
	if err != nil {
		return nil, err

	}
	// 再根据粉丝id列表获取粉丝信息列表
	return f.getUserInfo(followIdList)
}

// getUserInfo 根据id列表获取用户信息列表
func (f FollowService) getUserInfo(followIdList []uint) (list []model.UserInfo, err error) {
	// 再根据粉丝id列表获取粉丝信息列表
	for _, v := range followIdList {
		// 根据id获取用户信息
		users, err := f.UserDal.GetUserById(v)
		if err != nil {
			return nil, err
		}
		// 获取关注数
		followCount, err := f.FollowDal.GetFollowCount(v)
		if err != nil {
			return nil, err
		}
		// 获取粉丝数
		fansCount, err := f.FollowDal.GetFansCount(v)
		if err != nil {
			return nil, err
		}
		// 封装用户信息
		userInfo := model.UserInfo{
			ID:          users.ID,
			Username:    users.Username,
			Nickname:    users.Nickname,
			Avatar:      users.Avatar,
			FollowCount: followCount,
			FansCount:   fansCount,
		}
		list = append(list, userInfo)

	}
	// 	返回结果
	return list, nil
}

var _ abstract2.FollowService = (*FollowService)(nil)
