package v3

import (
	"errors"
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	"utopia-back/database/implement"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type FollowService struct {
	FollowDal abstract.FollowDal
	UserDal   abstract.UserDal
}

func NewFollowService() *FollowService {
	return &FollowService{
		FollowDal: &implement.FollowDal{},
		UserDal:   &implement.UserDal{},
	}
}

var ErrUserNotExist = errors.New("关注的用户不存在")

// Follow 关注
func (f FollowService) Follow(userId uint, followId uint) (err error) {
	// 判断followId是否存在
	_, err = f.UserDal.GetUserById(followId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotExist
		}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotExist
		}
		return err
	}
	// 取消关注
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
	for _, v := range fansIdList {
		userInfo, err := f.UserDal.GetUserInfoById(v)
		if err != nil {
			return nil, err
		}
		list = append(list, userInfo)
	}

	return list, nil
}

// GetFollowList 获取关注列表
func (f FollowService) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	//先获取自己关注的用户id列表
	followIdList, err := f.FollowDal.GetFollowIdList(userId)
	if err != nil {
		return nil, err

	}
	// 再根据粉丝id列表获取粉丝信息列表
	for _, v := range followIdList {
		userInfo, err := f.UserDal.GetUserInfoById(v)
		if err != nil {
			return nil, err
		}
		list = append(list, userInfo)
	}
	return list, nil
}

var _ abstract2.FollowService = (*FollowService)(nil)
