package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/database/implement"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type VideoService struct {
	VideoDal abstract.VideoDal
	UserDal  abstract.UserDal
}

func NewVideoService() *VideoService {
	return &VideoService{
		VideoDal: &implement.VideoDal{},
		UserDal:  &implement.UserDal{},
	}
}

func (v *VideoService) UpdateAvatar(uid uint, url string) error {
	return v.UserDal.UpdateAvatar(uid, url)
}

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

// 实现接口
var _ abstract2.VideoService = (*VideoService)(nil)
