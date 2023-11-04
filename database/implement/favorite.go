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

func (f FavoriteDal) GetFavoriteList(userId uint, lastTime uint, limitNum int) (list []uint, nextTime int, err error) {
	var favorites []*model.Favorite
	//返回的是videoId的列表
	res := f.Db.Where("user_id = ? AND updated_at > FROM_UNIXTIME(? / 1000) + INTERVAL (? % 1000) MICROSECOND AND status = 1", userId, lastTime, lastTime).
		Select("video_id,updated_at").
		Order("updated_at").
		Limit(limitNum).
		Find(&favorites)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	if len(favorites) == 0 {
		return
	}
	for _, v := range favorites {
		list = append(list, v.VideoID)
	}
	nextTime = int(favorites[len(favorites)-1].UpdatedAt.UnixMilli())
	return list, nextTime, err
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
