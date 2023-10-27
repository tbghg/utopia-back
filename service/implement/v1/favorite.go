package v1

import (
	"context"
	"strconv"
	"utopia-back/database/abstract"
	"utopia-back/pkg/redis"
	abstract2 "utopia-back/service/abstract"
)

type FavoriteService struct {
	Dal abstract.FavoriteDal
}

// 实现接口
var _ abstract2.FavoriteService = (*FavoriteService)(nil)

func (f FavoriteService) AddFavorite(userId uint, articleId uint) (err error) {
	return f.Dal.AddFavorite(userId, articleId)
}

func (f FavoriteService) DeleteFavorite(userId uint, video uint) (err error) {
	return f.Dal.DeleteFavorite(userId, video)
}

func (f FavoriteService) GetFavoriteList(userId uint) (list []uint, err error) {
	return f.Dal.GetFavoriteList(userId)
}

func (f FavoriteService) IsFavorite(userId uint, videoId uint) (isFavorite bool, err error) {
	//先去redis里面查找
	key := "favorite:" + "isFavorite:" + string(userId) + ":" + string(videoId)
	ctx := context.Background()
	value, err := redis.RDB.Get(ctx, key).Result()
	if err != nil {
		//redis里面没有，去数据库里面查找
		isFavorite, err = f.Dal.IsFavorite(userId, videoId)
		if err != nil {
			return false, err
		}
		//将结果存入redis
		redis.RDB.Set(ctx, key, isFavorite, 0)
		return isFavorite, nil
	}
	//redis里面有，直接返回
	res, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return res, nil

}

func (f FavoriteService) GetFavoriteCount(videoId uint) (count int64, err error) {
	//先去redis里面查找
	key := "favorite:count:" + string(videoId)
	ctx := context.Background()
	value, err := redis.RDB.Get(ctx, key).Result()
	if err != nil {
		//redis里面没有，去数据库里面查找
		count, err = f.Dal.GetFavoriteCount(videoId)
		if err != nil {
			return 0, err
		}
		//将结果存入redis
		redis.RDB.Set(ctx, key, count, 0)
		return count, nil
	}
	//redis里面有，直接返回
	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
