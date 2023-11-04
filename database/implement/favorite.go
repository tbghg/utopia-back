package implement

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type FavoriteDal struct {
	Db *gorm.DB
}

func (f FavoriteDal) AddFavorite(userId uint, videoId uint) (err error) {
	favorite := model.Favorite{
		UserID:  userId,
		VideoID: videoId,
		Status:  true,
	}
	// 存在则更新，不存在则插入
	res := f.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "userId"}, {Name: "videoId"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"status": favorite.Status}),
	}).Create(&favorite)
	return res.Error

}

func (f FavoriteDal) CancelFavorite(userId uint, video uint) (err error) {
	favorite := model.Favorite{
		UserID:  userId,
		VideoID: video,
		Status:  false,
	}
	res := f.Db.Model(&favorite).Where("user_id = ? AND video_id = ?", userId, video).Update("status", favorite.Status)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return err
}

func (f FavoriteDal) GetFavoriteList(userId uint) (list []uint, err error) {
	//返回的是videoId的列表
	res := f.Db.Where("user_id = ? AND status = 1", userId).Select("video_id").Find(&list)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return list, err
}

func (f FavoriteDal) IsFavorite(userId uint, videoId uint) (isFavorite bool, err error) {
	var favorite model.Favorite
	res := f.Db.Where("user_id = ? AND video_id = ? AND status = 1", userId, videoId).First(&favorite)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return favorite.Status, err
}

func (f FavoriteDal) GetFavoriteCount(videoId uint) (count int64, err error) {
	var favorite model.Favorite
	res := f.Db.Where("video_id = ? AND status = 1", videoId).Find(&favorite).Count(&count)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return count, err
}

var _ abstract.FavoriteDal = (*FavoriteDal)(nil)
