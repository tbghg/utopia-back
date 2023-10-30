package abstract

import "utopia-back/model"

type VideoDal interface {
	// CreateVideo 创建视频
	CreateVideo(video *model.Video) (id uint, err error)
	// IsVideoExist 判断视频是否存在
	IsVideoExist(videoId uint) (err error)
	// GetVideoByType 查找某分区下的视频
	GetVideoByType(lastTime string, videoTypeId uint) (video []*model.Video, err error)
}
