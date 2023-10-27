package abstract

type FavoriteService interface {
	// AddFavorite 添加收藏
	AddFavorite(userId uint, articleId uint) (err error)
	// DeleteFavorite 删除收藏
	DeleteFavorite(userId uint, video uint) (err error)
	// GetFavoriteList 获取收藏列表
	GetFavoriteList(userId uint) (list []uint, err error)
	// IsFavorite 是否收藏
	IsFavorite(userId uint, videoId uint) (isFavorite bool, err error)
	// GetFavoriteCount 获取收藏数
	GetFavoriteCount(videoId uint) (count int64, err error)
}
