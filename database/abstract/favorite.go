package abstract

type FavoriteDal interface {
	// AddFavorite 添加收藏
	AddFavorite(userId uint, videoId uint) (err error)
	// CancelFavorite  取消收藏
	CancelFavorite(userId uint, video uint) (err error)
	// GetFavoriteList 获取收藏列表
	GetFavoriteList(userId uint, lastTime uint, limitNum int) (list []uint, nextTime int, err error)
	// IsFavorite 是否收藏
	IsFavorite(userId uint, videoId uint) (isFavorite bool, err error)
	// GetFavoriteCount 获取收藏数
	GetFavoriteCount(videoId uint) (count int64, err error)
}
