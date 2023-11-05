package abstract

import "utopia-back/model"

type VideoDal interface {
	// CreateVideo 创建视频
	CreateVideo(video *model.Video) (id uint, err error)
	// IsVideoExist 判断视频是否存在
	IsVideoExist(videoId uint) (err error)
	// GetVideoByType 查找某分区下的视频
	GetVideoByType(lastTime uint, videoTypeId uint, limitNum int) (videos []*model.Video, err error)
	// GetUploadVideos 查找某分区下的视频
	GetUploadVideos(lastTime uint, uid uint, limitNum int) (videos []*model.Video, err error)
	// GetPopularVideos 获取热门视频
	GetPopularVideos(limitNum int) (videoIds []*model.VideoCount, err error)
	// GetVideoInfoById 获取视频信息
	GetVideoInfoById(videoIds []uint) (videos []*model.Video, err error)
	// SearchVideos 查找视频
	SearchVideos(search string, limitNum int) (videos []*model.Video, err error)
}
