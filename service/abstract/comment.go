package abstract

import "utopia-back/model"

type CommentService interface {
	// Comment 评论
	Comment(userId uint, videoId uint, content string) (err error)
	// CommentList 评论列表
	CommentList(videoId uint) (comments []*model.CommentInfo, err error)
}
