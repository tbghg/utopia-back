package v1

import (
	"utopia-back/cache"
	"utopia-back/database/abstract"
	abstract2 "utopia-back/service/abstract"
)

// 实现接口
var _ abstract2.LikeService = (*LikeService)(nil)

type LikeService struct {
	LikeDal  abstract.LikeDal
	VideoDal abstract.VideoDal
}

func (l LikeService) Like(userId uint, videoId uint) (err error) {
	//判断videoId是否存在
	err = l.VideoDal.IsVideoExist(videoId)
	if err != nil {
		return err
	}
	//更新数据库
	_, err = l.LikeDal.Like(userId, videoId)
	if err != nil {
		return err
	}
	// 更新缓存
	key := cache.VideoLikeCountKey(videoId)
	err = cache.RDB.Incr(cache.Ctx, key).Err()

	if err != nil {
		return err
	}
	return nil
}

func (l LikeService) UnLike(userId uint, videoId uint) (err error) {
	//判断videoId是否存在
	err = l.VideoDal.IsVideoExist(videoId)
	if err != nil {
		return err
	}

	//更新数据库
	_, err = l.LikeDal.UnLike(userId, videoId)
	if err != nil {
		return err
	}
	// 更新缓存
	key := cache.VideoLikeCountKey(videoId)
	err = cache.RDB.Decr(cache.Ctx, key).Err()
	if err != nil {
		return err
	}
	return err
}

func (l LikeService) IsLike(userId uint, videoId uint) (isLike bool, err error) {
	return l.LikeDal.IsLike(userId, videoId)
}

func (l LikeService) GetLikeCount(videoId uint) (count int64, err error) {
	// 构造key
	key := cache.VideoLikeCountKey(videoId)
	// 缓存层取数据
	res, err := cache.GetStringCache(func() (interface{}, error) { return l.LikeDal.GetLikeCount(videoId) }, key, cache.TypeInt64)
	// 返回结果
	if err != nil {
		return 0, err
	}

	return res.(int64), nil
}
