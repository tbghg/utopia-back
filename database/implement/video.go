package implement

import (
	"utopia-back/database"
	"utopia-back/model"
)

type VideoDal struct {
}

func (v *VideoDal) CreateVideo(video *model.Video) (id uint, err error) {
	res := database.DB.Create(&video)
	if res.Error != nil {
		return 0, res.Error
	}
	return video.ID, nil
}

func (v *VideoDal) IsVideoExist(videoId uint) (exist bool, err error) {
	res := database.DB.Where("id = ?", videoId).First(&model.Video{})
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}
