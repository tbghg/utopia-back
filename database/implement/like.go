package implement

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type LikeDal struct{ Db *gorm.DB }

func (l *LikeDal) Like(userId uint, videoId uint) (err error) {
	res := l.Db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "video_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"status": true}),
		}).Create(&model.Like{
		UserID:  userId,
		VideoID: videoId,
		Status:  true,
	})
	return res.Error
}

func (l *LikeDal) UnLike(userId uint, videoId uint) (err error) {
	res := l.Db.Model(&model.Like{}).Where("user_id = ? AND video_id = ?", userId, videoId).Update("status", false)
	return res.Error
}

func (l *LikeDal) IsLike(userId uint, videoId uint) (isLike bool, err error) {
	var like model.Like
	res := l.Db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&like)
	return like.Status, res.Error
}

func (l *LikeDal) GetLikeCount(videoId uint) (count int64, err error) {
	var likeCount int64
	res := l.Db.Model(&model.Like{}).Where("video_id = ?", videoId).Count(&likeCount)
	return likeCount, res.Error
}

var _ abstract.LikeDal = (*LikeDal)(nil)
