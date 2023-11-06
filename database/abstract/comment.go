package abstract

import "utopia-back/model"

type CommentDal interface {
	// Comment 评论
	Comment(userId uint, videoId uint, content string) (err error)
	// CommentList 评论列表
	CommentList(videoId uint, lastTime uint, limitNum int) (comments []*model.CommentInfo, err error)
	// CommentNum 评论数
	CommentNum(videoId uint) (commentNum int, err error)
}
