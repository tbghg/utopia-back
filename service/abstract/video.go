package abstract

import "utopia-back/model"

type VideoService interface {
	// GetCategoryVideos 获取某分区下的视频
	GetCategoryVideos(uid uint, lastTime string, videoTypeId uint) ([]*model.VideoInfo, string, error)
}
