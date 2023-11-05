package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type CommentService struct {
	CommentDal abstract.CommentDal
}

func (c CommentService) Comment(userId uint, videoId uint, content string) (err error) {
	return c.CommentDal.Comment(userId, videoId, content)
}

func (c CommentService) CommentList(videoId uint) (comments []*model.CommentInfo, err error) {
	return c.CommentDal.CommentList(videoId)
}

// 实现接口
var _ abstract2.CommentService = (*CommentService)(nil)
