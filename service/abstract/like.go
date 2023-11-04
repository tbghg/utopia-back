package abstract

type LikeService interface {
	// Like 点赞
	Like(userId uint, videoId uint) (err error)
	// UnLike 取消点赞
	UnLike(userId uint, videoId uint) (err error)
	// GetLikeCount 获取点赞数
	//GetLikeCount(videoId uint) (count int64, err error)
	//GetBatchIsLike 批量获取是否为视频点赞
	//GetBatchIsLike(userId uint, videoId []uint) (isLike []bool, err error)
	//IsLike 是否点赞
	//IsLike(userId uint, videoId uint) (isLike bool, err error)
}
