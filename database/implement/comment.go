package implement

import (
	"errors"
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type CommentDal struct{ Db *gorm.DB }

func (c CommentDal) Comment(userId uint, videoId uint, content string) (err error) {
	comment := model.Comment{
		VideoID: videoId,
		UserID:  userId,
		Content: content,
	}
	return c.Db.Model(model.Comment{}).Create(&comment).Error
}

func (c CommentDal) CommentList(videoId uint) (commentInfo []*model.CommentInfo, err error) {

	err = c.Db.Model(model.Comment{}).
		Select("comments.content, users.nickname, users.avatar").
		Joins("JOIN users ON comments.user_id = users.id").
		Where("comments.video_id = ?", videoId).
		Find(&commentInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return commentInfo, err
}

var _ abstract.CommentDal = (*CommentDal)(nil)
