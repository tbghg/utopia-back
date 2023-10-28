package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/database/implement"
	abstract2 "utopia-back/service/abstract"
)

type FavoriteService struct {
	FavoriteDal abstract.FavoriteDal
	VideoDal    abstract.VideoDal
}

func NewFavoriteService() *FavoriteService {
	return &FavoriteService{
		FavoriteDal: &implement.FavoriteDal{},
	}
}

// 实现接口
var _ abstract2.FavoriteService = (*FavoriteService)(nil)

// AddFavorite 添加收藏
func (f FavoriteService) AddFavorite(userId uint, videoId uint) (err error) {
	// 判断video是否存在
	exist, err := f.VideoDal.IsVideoExist(videoId)
	if err != nil || !exist {
		return err
	}
	// 添加收藏
	return f.FavoriteDal.AddFavorite(userId, videoId)
}

// CancelFavorite  取消收藏
func (f FavoriteService) CancelFavorite(userId uint, videoId uint) (err error) {
	// 判断video是否存在
	exist, err := f.VideoDal.IsVideoExist(videoId)
	if err != nil || !exist {
		return err
	}
	// 取消收藏
	return f.FavoriteDal.CancelFavorite(userId, videoId)
}
