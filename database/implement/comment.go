package implement

import (
	"errors"
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type CommentDal struct{ Db *gorm.DB }

func (c CommentDal) CommentNum(videoId uint) (commentNum int, err error) {
	var count int64
	res := c.Db.Model(model.Comment{}).Where("video_id = ?", videoId).Count(&count)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		err = res.Error
	}
	commentNum = int(count)
	return
}

func (c CommentDal) Comment(userId uint, videoId uint, content string) (err error) {
	comment := model.Comment{
		VideoID: videoId,
		UserID:  userId,
		Content: content,
	}
	return c.Db.Model(model.Comment{}).Create(&comment).Error
}

func (c CommentDal) CommentList(videoId uint, lastTime uint, limitNum int) (commentInfo []*model.CommentInfo, err error) {

	err = c.Db.Model(model.Comment{}).
		Select("comments.content, comments.updated_at, users.nickname, users.avatar").
		Joins("JOIN users ON comments.user_id = users.id").
		Where("comments.video_id = ?", videoId).
		Where("comments.updated_at > FROM_UNIXTIME(? / 1000) + INTERVAL (? % 1000) MICROSECOND", lastTime, lastTime).
		Limit(limitNum).
		Find(&commentInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return commentInfo, err
}

var _ abstract.CommentDal = (*CommentDal)(nil)
