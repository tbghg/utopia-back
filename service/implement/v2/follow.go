package v2

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type FollowService struct {
	FollowDal abstract.FollowDal
	UserDal   abstract.UserDal
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
	// 根据粉丝id列表获取粉丝信息列表
	return f.getUserInfoList(fansIdList)
}

// GetFollowList 获取关注列表
func (f FollowService) GetFollowList(userId uint) (list []model.UserInfo, err error) {
	//先获取关注id列表
	followIdList, err := f.FollowDal.GetFollowIdList(userId)
	if err != nil {
		return nil, err
	}
	// 根据关注id列表获取关注信息列表
	return f.getUserInfoList(followIdList)

}

func (f FollowService) getUserInfoList(fansIdList []uint) (list []model.UserInfo, err error) {
	// 容器
	userMap := make(map[uint]model.User)
	fansCountMap := make(map[uint]int64)
	followCountMap := make(map[uint]int64)
	// 通道
	userMapChan := make(chan map[uint]model.User)
	fansCountMapChan := make(chan map[uint]int64)
	followCountMapChan := make(chan map[uint]int64)
	errChan := make(chan error)

	// 根据用户id列表的获取用户信息
	go func() {
		for _, fansId := range fansIdList {
			user, err := f.UserDal.GetUserById(fansId)
			if err != nil {
				errChan <- errors.New(fmt.Sprintf("获取该用户信息失败, 用户id：%v, 错误原因：%v", fansId, err))
				return
			}
			userMap[fansId] = user
		}
		userMapChan <- userMap
	}()

	// 根据用户id列表的获取关注数
	go func() {
		for _, fansId := range fansIdList {
			followCount, err := f.FollowDal.GetFollowCount(fansId)
			if err != nil {
				errChan <- errors.New(fmt.Sprintf("获取该用户关注数失败, 用户id：%v, 错误原因：%v", fansId, err))
				return
			}
			followCountMap[fansId] = followCount
		}
		followCountMapChan <- followCountMap
	}()
	// 根据用户id列表的获取粉丝数
	go func() {
		for _, fansId := range fansIdList {
			fansCount, err := f.FollowDal.GetFansCount(fansId)
			if err != nil {
				errChan <- errors.New(fmt.Sprintf("获取该用户粉丝数失败, 用户id：%v, 错误原因：%v", fansId, err))
				return
			}
			fansCountMap[fansId] = fansCount
		}
		fansCountMapChan <- fansCountMap
	}()
	// 从通道中获取数据
	for i := 0; i < 3; i++ {
		select {
		case userMap = <-userMapChan:
		case fansCountMap = <-fansCountMapChan:
		case followCountMap = <-followCountMapChan:
		case err = <-errChan:
			return nil, err
		}
	}

	// 将粉丝信息列表、粉丝的关注数、粉丝的粉丝数合并到一起
	var userInfoMap = make(map[uint]model.UserInfo)
	for _, id := range fansIdList {
		userInfoMap[id] = model.UserInfo{
			ID:          id,
			Avatar:      userMap[id].Avatar,
			Nickname:    userMap[id].Nickname,
			Username:    userMap[id].Username,
			FollowCount: followCountMap[id],
			FansCount:   fansCountMap[id],
		}
		list = append(list, userInfoMap[id])
	}
	return list, nil
}
