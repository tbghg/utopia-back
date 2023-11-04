package abstract

type LikeDal interface {
	// Like 点赞
	Like(userId uint, videoId uint) (rowsAffected int, err error)
	// UnLike 取消点赞
	UnLike(userId uint, videoId uint) (rowsAffected int, err error)
	// IsLike 是否点赞
	IsLike(userId uint, videoId uint) (isLike bool, err error)
	// GetLikeCount 获取点赞数
	GetLikeCount(videoId uint) (count int64, err error)
	// GetLikeUserId 获取点赞用户
	GetLikeUserId(videoId uint) (user []uint, err error)
	// GetUserLikedVideosWithLimit 获取用户点过赞的视频id(前itemNum条)
	GetUserLikedVideosWithLimit(userId uint, itemNum int) (videoId []uint, err error)
	// BatchIsLike 批量判断是否点赞
	BatchIsLike(userId uint, videoId []uint) (videoIds []uint, err error)
	// UpdateLikeCount 更新视频点赞数
	UpdateLikeCount(videoId uint, num int)
}
