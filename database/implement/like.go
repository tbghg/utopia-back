package implement

import (
	"gorm.io/gorm/clause"
	"utopia-back/database"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type LikeDal struct{}

func (l *LikeDal) Like(userId uint, videoId uint) (err error) {
	res := database.DB.Clauses(
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
	res := database.DB.Model(&model.Like{}).Where("user_id = ? AND video_id = ?", userId, videoId).Update("status", false)
	return res.Error
}

func (l *LikeDal) IsLike(userId uint, videoId uint) (isLike bool, err error) {
	var like model.Like
	res := database.DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&like)
	return like.Status, res.Error
}

func (l *LikeDal) GetLikeCount(videoId uint) (count int64, err error) {
	var likeCount int64
	res := database.DB.Model(&model.Like{}).Where("video_id = ?", videoId).Count(&likeCount)
	return likeCount, res.Error
}

var _ abstract.LikeDal = (*LikeDal)(nil)
