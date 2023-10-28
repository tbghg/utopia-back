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
