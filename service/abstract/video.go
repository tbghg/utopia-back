package abstract

import "utopia-back/model"

type VideoService interface {
	// GetCategoryVideos 获取某分区下的视频
	GetCategoryVideos(uid uint, lastTime uint, videoTypeId uint) ([]*model.VideoInfo, int, error)
	// GetPopularVideos 获取热门视频
	GetPopularVideos(uid uint, version int, score float64) (videoInfo []*model.VideoInfo, nextScore float64, nextVersion int, err error)
	// GetRecommendVideos 获取推荐视频
	GetRecommendVideos(uid uint, lastTime uint) ([]*model.VideoInfo, int, error)
	// GetFavoriteVideos 收藏视频列表
	GetFavoriteVideos(uid uint, targetUid uint, lastTime uint) (videoInfos []*model.VideoInfo, nextTime int, err error)
	// GetUploadVideos 发布视频列表
	GetUploadVideos(uid uint, targetUid uint, lastTime uint) ([]*model.VideoInfo, int, error)
	// SearchVideos 搜索
	SearchVideos(uid uint, search string) ([]*model.VideoInfo, error)
}
