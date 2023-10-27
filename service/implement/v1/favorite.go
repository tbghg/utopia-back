package v1

import (
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
	// 构造key
	key := "favorite:" + "isFavorite:" + string(userId) + ":" + string(videoId)
	// 缓存
	res, err := redis.Cache(func() (interface{}, error) { return f.Dal.IsFavorite(userId, videoId) }, key, redis.TypeBool)
	// 返回结果
	if err != nil {
		return false, err
	}
	return res.(bool), nil

}

func (f FavoriteService) GetFavoriteCount(videoId uint) (count int64, err error) {
	// 构造key
	key := "favorite:count:" + string(videoId)
	// 缓存层取数据
	res, err := redis.Cache(func() (interface{}, error) { return f.Dal.GetFavoriteCount(videoId) }, key, redis.TypeInt64)
	// 返回结果
	if err != nil {
		return 0, err
	}
	return res.(int64), nil
}
