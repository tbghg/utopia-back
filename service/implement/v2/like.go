package v2

import (
	"errors"
	"strconv"
	"time"
	"utopia-back/database/abstract"
	"utopia-back/pkg/redis"
	abstract2 "utopia-back/service/abstract"
)

type LikeService struct {
	LikeDal abstract.LikeDal
}

var kerPrefix = "v2:like:"
var originExpireTime = 24 * time.Hour

//todo 布隆过滤器

func (l LikeService) Like(userId uint, videoId uint) (err error) {
	key := kerPrefix + strconv.Itoa(int(videoId))
	// 直接操作redis Set重复的元素不会被添加
	err = redis.RDB.SAdd(redis.Ctx, key, userId).Err()
	if err != nil {
		return err
	}
	// 设置过期时间
	return redis.RDB.Expire(redis.Ctx, key, originExpireTime).Err()

}

func (l LikeService) UnLike(userId uint, videoId uint) (err error) {
	key := kerPrefix + strconv.Itoa(int(videoId))
	err = redis.RDB.SRem(redis.Ctx, key, userId).Err()
	if err != nil {
		return err
	}
	// 设置过期时间
	return redis.RDB.Expire(redis.Ctx, key, originExpireTime).Err()
}

const (
	modeLikeCount = iota
	modeIsLike
)

func (l LikeService) IsLike(userId uint, videoId uint) (isLike bool, err error) {
	res, err := l.cache(videoId, userId, modeIsLike)
	isLike = res.(bool)
	return isLike, err
}

func (l LikeService) GetLikeCount(videoId uint) (count int64, err error) {
	res, err := l.cache(videoId, 0, modeLikeCount)
	count = res.(int64)
	return count, err
}

func (l LikeService) cache(videoId, userId uint, mode int) (res interface{}, err error) {
	key := kerPrefix + strconv.Itoa(int(videoId))
	//先判断缓存中Key是否存在
	exist, err := redis.RDB.Exists(redis.Ctx, key).Result()
	//Redis发生错误或者缓存不存在
	if err != nil || exist == 0 {
		//从数据库中获取
		userIDs, err := l.LikeDal.GetLikeUserId(videoId)
		//数据库发生错误或者不存在 直接返回
		if err != nil || len(userIDs) == 0 {
			return res, err
		}
		//更新缓存

		if err = redis.RDB.SAdd(redis.Ctx, key, userIDs).Err(); err != nil {
			return res, err
		}

		//设置过期时间
		if err = redis.RDB.Expire(redis.Ctx, key, originExpireTime).Err(); err != nil {
			return res, err
		}
		//返回结果
		switch mode {
		case modeLikeCount:

			return redis.RDB.SCard(redis.Ctx, key).Result()
		case modeIsLike:
			return redis.RDB.SIsMember(redis.Ctx, key, userIDs).Result()
		default:
			return nil, errors.New("模式错误")
		}

	}
	//缓存中存在
	switch mode {
	case modeLikeCount:
		res, err = redis.RDB.SCard(redis.Ctx, key).Result()
	case modeIsLike:
		res, err = redis.RDB.SIsMember(redis.Ctx, key, userId).Result()
	default:
		return nil, errors.New("模式错误")
	}

	//判断过期时间 是否超过生命周期2/3
	ttl, err := redis.RDB.TTL(redis.Ctx, key).Result()
	if err != nil {
		return res, err
	}
	if ttl < originExpireTime/3*2 {
		//	延长生命周期
		redis.RDB.Expire(redis.Ctx, key, originExpireTime)
	}

	return res, nil
}

// 实现接口
var _ abstract2.LikeService = (*LikeService)(nil)
