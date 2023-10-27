package abstract

import "utopia-back/model"

type VideoDal interface {
	CreateVideo(video *model.Video) (id uint, err error)
}
