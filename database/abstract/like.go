package abstract

type LikeDal interface {
	// Like 点赞
	Like(userId uint, videoId uint) (err error)
	// UnLike 取消点赞
	UnLike(userId uint, videoId uint) (err error)
	// IsLike 是否点赞
	IsLike(userId uint, videoId uint) (isLike bool, err error)
	// GetLikeCount 获取点赞数
	GetLikeCount(videoId uint) (count int64, err error)
	// GetLikeUserId 获取点赞用户
	GetLikeUserId(videoId uint) (user []uint, err error)
}
