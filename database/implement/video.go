package implement

import (
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type VideoDal struct {
	Db *gorm.DB
}

func (v *VideoDal) GetVideoInfoById(videoIds []uint) (videos []*model.Video, err error) {
	res := v.Db.Model(model.Video{}).
		Where("id IN ?", videoIds).
		Find(&videos)
	err = res.Error

	// 将视频信息与videoIds关联，保持原始顺序
	videoMap := make(map[uint]*model.Video)
	for _, video := range videos {
		videoMap[video.ID] = video
	}

	// 按照videoIds的顺序构建排序后的视频切片
	sortedVideos := make([]*model.Video, len(videoIds))
	for i, id := range videoIds {
		sortedVideos[i] = videoMap[id]
	}

	return sortedVideos, err
}

func (v *VideoDal) GetUploadVideos(lastTime uint, uid uint, limitNum int) (videos []*model.Video, err error) {
	res := v.Db.Model(model.Video{}).
		Where("created_at >  FROM_UNIXTIME(? / 1000) + INTERVAL (? % 1000) MICROSECOND AND author_id = ?", lastTime, lastTime, uid).
		Order("created_at").Limit(limitNum).Find(&videos)
	err = res.Error
	return
}

func (v *VideoDal) GetVideoByType(lastTime uint, videoTypeId uint, limitNum int) (videos []*model.Video, err error) {
	res := v.Db.Model(model.Video{}).
		Where("created_at >  FROM_UNIXTIME(? / 1000) + INTERVAL (? % 1000) MICROSECOND  and video_type_id = ?", lastTime, lastTime, videoTypeId).
		Order("created_at").Limit(limitNum).Find(&videos)
	err = res.Error
	return
}

func (v *VideoDal) CreateVideo(video *model.Video) (id uint, err error) {
	res := v.Db.Create(&video)
	if res.Error != nil {
		return 0, res.Error
	}
	resLC := v.Db.Create(&model.LikeCount{
		Count:   0,
		VideoID: video.ID,
	})
	if res.Error != nil {
		return 0, resLC.Error
	}
	return video.ID, nil
}

func (v *VideoDal) IsVideoExist(videoId uint) (err error) {
	res := v.Db.Where("id = ?", videoId).First(&model.Video{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

var _ abstract.VideoDal = (*VideoDal)(nil)
