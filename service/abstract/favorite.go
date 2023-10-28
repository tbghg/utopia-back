package abstract

type FavoriteService interface {
	// AddFavorite 添加收藏
	AddFavorite(userId uint, articleId uint) (err error)
	// CancelFavorite 取消收藏
	CancelFavorite(userId uint, video uint) (err error)
}
