package v1

import (
	"strconv"
	"utopia-back/database/abstract"
	"utopia-back/database/implement"
	"utopia-back/pkg/redis"
	abstract2 "utopia-back/service/abstract"
)

// 实现接口
var _ abstract2.LikeService = (*LikeService)(nil)

type LikeService struct {
	LikeDal  abstract.LikeDal
	VideoDal abstract.VideoDal
}

func NewLikeService() *LikeService {
	return &LikeService{
		LikeDal: &implement.LikeDal{},
	}
}

func (l LikeService) Like(userId uint, videoId uint) (err error) {
	//判断videoId是否存在
	err = l.VideoDal.IsVideoExist(videoId)
	if err != nil {
		return err
	}
	//更新数据库
	err = l.LikeDal.Like(userId, videoId)
	if err != nil {
		return err
	}
	// 更新缓存
	err = redis.RDB.Incr(redis.Ctx, "like:count:"+strconv.Itoa(int(videoId))).Err()
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
	err = l.LikeDal.UnLike(userId, videoId)
	if err != nil {
		return err
	}
	// 更新缓存
	err = redis.RDB.Decr(redis.Ctx, "like:count:"+strconv.Itoa(int(videoId))).Err()
	if err != nil {
		return err
	}
	return l.LikeDal.UnLike(userId, videoId)
}

func (l LikeService) IsLike(userId uint, videoId uint) (isLike bool, err error) {
	return l.LikeDal.IsLike(userId, videoId)
}

func (l LikeService) GetLikeCount(videoId uint) (count int64, err error) {
	// 构造key
	key := "like:count:" + strconv.Itoa(int(videoId))
	// 缓存层取数据
	res, err := redis.GetStringCache(func() (interface{}, error) { return l.LikeDal.GetLikeCount(videoId) }, key, redis.TypeInt64)
	// 返回结果
	if err != nil {
		return 0, err
	}

	return res.(int64), nil
}
