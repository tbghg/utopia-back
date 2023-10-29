package v1

import (
	"errors"
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	abstract2 "utopia-back/service/abstract"
)

type FavoriteService struct {
	FavoriteDal abstract.FavoriteDal
	VideoDal    abstract.VideoDal
}

//func NewFavoriteService() *FavoriteService {
//	return &FavoriteService{
//		FavoriteDal: &implement.FavoriteDal{},
//		VideoDal:    &implement.VideoDal{},
//	}
//}

// 实现接口
var _ abstract2.FavoriteService = (*FavoriteService)(nil)

var ErrVideoNotExist = errors.New("视频不存在")

// AddFavorite 添加收藏
func (f FavoriteService) AddFavorite(userId uint, videoId uint) (err error) {
	// 判断video是否存在
	err = f.VideoDal.IsVideoExist(videoId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrVideoNotExist
		}
		return err
	}
	// 添加收藏
	return f.FavoriteDal.AddFavorite(userId, videoId)
}

// CancelFavorite  取消收藏
func (f FavoriteService) CancelFavorite(userId uint, videoId uint) (err error) {
	// 判断video是否存在
	err = f.VideoDal.IsVideoExist(videoId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrVideoNotExist
		}
		return err
	}
	// 取消收藏
	return f.FavoriteDal.CancelFavorite(userId, videoId)
}
