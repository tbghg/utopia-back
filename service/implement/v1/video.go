package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type VideoService struct {
	Dal abstract.VideoDal
}

// 实现接口
var _ abstract2.VideoService = (*VideoService)(nil)

func (v *VideoService) UploadVideoCallback(authorId uint, url string, coverUrl string, describe string) (err error) {
	video := &model.Video{
		AuthorID: authorId,
		PlayUrl:  url,
		CoverUrl: coverUrl,
		Describe: describe,
	}
	_, err = v.Dal.CreateVideo(video)
	if err != nil {
		return err
	}
	return
}
