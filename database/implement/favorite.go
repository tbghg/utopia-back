package implement

import (
	"utopia-back/database"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type FavoriteDal struct{}

func (f FavoriteDal) AddFavorite(userId uint, videoId uint) (err error) {
	res := database.DB.Create(&model.Favorite{
		UserID:  userId,
		VideoID: videoId,
	})
	return res.Error

}

func (f FavoriteDal) DeleteFavorite(userId uint, video uint) (err error) {
	res := database.DB.Where("user_id = ? AND video_id = ?", userId, video).Delete(&model.Favorite{})
	return res.Error
}

func (f FavoriteDal) GetFavoriteList(userId uint) (list []uint, err error) {
	//返回的是videoId的列表
	res := database.DB.Where("user_id = ?", userId).Select("video_id").Find(&list)
	return list, res.Error
}

func (f FavoriteDal) IsFavorite(userId uint, videoId uint) (isFavorite bool, err error) {
	var favorite model.Favorite
	res := database.DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (f FavoriteDal) GetFavoriteCount(videoId uint) (count int64, err error) {
	var favorite model.Favorite
	res := database.DB.Where("video_id = ?", videoId).Find(&favorite).Count(&count)
	return count, res.Error
}

var _ abstract.FavoriteDal = (*FavoriteDal)(nil)
