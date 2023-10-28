package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/database/implement"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type VideoService struct {
	VideoDal abstract.VideoDal
}

func NewVideoService() *VideoService {
	return &VideoService{
		VideoDal: &implement.VideoImpl{},
	}
}

// 实现接口
var _ abstract2.VideoService = (*VideoService)(nil)

func (v *VideoService) UploadVideoCallback(authorId uint, url string, coverUrl string, describe string, videoType uint) (err error) {
	video := &model.Video{
		AuthorID:  authorId,
		PlayUrl:   url,
		CoverUrl:  coverUrl,
		VideoType: videoType,
		Describe:  describe,
	}
	_, err = v.VideoDal.CreateVideo(video)
	if err != nil {
		return err
	}
	return
}
