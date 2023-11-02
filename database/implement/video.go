package implement

import (
	"gorm.io/gorm"
	"utopia-back/database/abstract"
	"utopia-back/model"
)

type VideoDal struct {
	Db *gorm.DB
}

func (v *VideoDal) GetVideoByType(lastTime string, videoTypeId uint) (videos []*model.Video, err error) {
	res := v.Db.Model(model.Video{}).
		Where("created_at > from_unixtime(?) and video_type_id = ?", lastTime, videoTypeId).
		Order("created_at").Limit(3).Find(&videos)
	err = res.Error
	return
}

func (v *VideoDal) CreateVideo(video *model.Video) (id uint, err error) {
	res := v.Db.Create(&video)
	if res.Error != nil {
		return 0, res.Error
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
