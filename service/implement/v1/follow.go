package v1

import (
	"errors"
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
	}
}

func (f FollowService) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	return f.FollowDal.GetFollowList(userId)
}

var ErrUserNotExist = errors.New("用户不存在")

// Follow 关注
func (f FollowService) Follow(userId uint, followId uint) (err error) {
	// 判断followId是否存在
	user, err := f.UserDal.GetUserById(followId)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return ErrUserNotExist
	}
	// 关注
	return f.FollowDal.Follow(userId, followId)
}

// UnFollow 取消关注
func (f FollowService) UnFollow(userId uint, followId uint) (err error) {
	// 判断followId是否存在
	user, err := f.UserDal.GetUserById(followId)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return ErrUserNotExist
	}
	// 取消关注
	return f.FollowDal.UnFollow(userId, followId)
}

// GetFansList 获取粉丝列表
func (f FollowService) GetFansList(userId uint) (list []model.UserInfo, err error) {
	return f.FollowDal.GetFansList(userId)
}

var _ abstract2.FollowService = (*FollowService)(nil)
