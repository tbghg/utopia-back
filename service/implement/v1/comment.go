package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

const commentListLimitNum = 20

type CommentService struct {
	CommentDal abstract.CommentDal
}

func (c CommentService) Comment(userId uint, videoId uint, content string) (err error) {
	return c.CommentDal.Comment(userId, videoId, content)
}

func (c CommentService) CommentList(videoId uint, lastTime uint) (comments []*model.CommentInfo, nextTime int, err error) {

	comments, err = c.CommentDal.CommentList(videoId, lastTime, commentListLimitNum)
	if err != nil || len(comments) == 0 {
		return nil, -1, err
	}
	nextTime = int(comments[len(comments)-1].UpdatedAt.UnixMilli())
	return
}

// 实现接口
var _ abstract2.CommentService = (*CommentService)(nil)
