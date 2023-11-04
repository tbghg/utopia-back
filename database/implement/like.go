package implement

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type LikeDal struct{ Db *gorm.DB }

// BatchIsLike 批量查询是否点过赞
func (l *LikeDal) BatchIsLike(userId uint, videoIds []uint) (likedVideoIds []uint, err error) {

	var builder strings.Builder
	for i, v := range videoIds {
		builder.WriteString(strconv.Itoa(int(v)))
		if i != len(videoIds)-1 {
			builder.WriteString(",")
		}
	}
	fields := builder.String()

	res := l.Db.Model(&model.Like{}).
		Select("video_id").
		Where("user_id = ? and status = 1 and video_id IN (?)", userId, fields).
		Find(&likedVideoIds)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return likedVideoIds, err
}

func (l *LikeDal) GetUserLikedVideosWithLimit(userId uint, itemNum int) (videoId []uint, err error) {
	res := l.Db.Model(&model.Like{}).Select("video_id").
		Where("user_id = ? AND status = ?", userId, true).
		Order("id desc").
		Limit(itemNum).Find(&videoId)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return videoId, err
}

func (l *LikeDal) GetLikeUserId(videoId uint) (user []uint, err error) {
	res := l.Db.Model(&model.Like{}).Where("video_id = ? AND status = ?", videoId, true).Find(&user)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return user, err
}

func (l *LikeDal) Like(userId uint, videoId uint) (rowsAffected int, err error) {
	res := l.Db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "video_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"status": true}),
		}).Create(&model.Like{
		UserID:  userId,
		VideoID: videoId,
		Status:  true,
	})
	return int(res.RowsAffected), res.Error
}

func (l *LikeDal) UpdateLikeCount(videoId uint, num int) {
	if num > 0 {
		l.Db.Model(&model.LikeCount{}).
			Where("video_id = ?", videoId).
			UpdateColumn("count", gorm.Expr("count + ?", num))
	} else if num < 0 {
		l.Db.Model(&model.LikeCount{}).
			Where("video_id = ?", videoId).
			UpdateColumn("count", gorm.Expr("count - ?", -1*num))
	}
}

func (l *LikeDal) UnLike(userId uint, videoId uint) (rowsAffected int, err error) {
	res := l.Db.Model(&model.Like{}).
		Where("user_id = ? AND video_id = ? AND status = 1", userId, videoId).
		Update("status", false)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return int(res.RowsAffected), err
}

func (l *LikeDal) IsLike(userId uint, videoId uint) (isLike bool, err error) {
	var like model.Like
	res := l.Db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&like)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return like.Status, err
}

func (l *LikeDal) GetLikeCount(videoId uint) (count int64, err error) {
	var likeCount int64
	res := l.Db.Model(&model.Like{}).Where("video_id = ?", videoId).Count(&likeCount)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	return likeCount, err
}

var _ abstract.LikeDal = (*LikeDal)(nil)
