package abstract

import "utopia-back/model"

type VideoDal interface {
	// CreateVideo 创建视频
	CreateVideo(video *model.Video) (id uint, err error)
	// IsVideoExist 判断视频是否存在
	IsVideoExist(videoId uint) (exist bool, err error)
}
