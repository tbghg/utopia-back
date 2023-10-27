package implement

import (
	"utopia-back/database"
	"utopia-back/model"
)

type VideoImpl struct {
}

func (v *VideoImpl) CreateVideo(video *model.Video) (id uint, err error) {
	res := database.DB.Create(&video)
	if res.Error != nil {
		return 0, res.Error
	}
	return video.ID, nil
}
