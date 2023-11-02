package abstract

import "utopia-back/model"

type VideoService interface {
	// GetCategoryVideos 获取某分区下的视频
	GetCategoryVideos(uid uint, lastTime uint, videoTypeId uint) ([]*model.VideoInfo, int, error)
}
