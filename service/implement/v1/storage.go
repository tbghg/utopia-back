package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type StorageService struct {
	VideoDal abstract.VideoDal
	UserDal  abstract.UserDal
}

//func NewStorageService() *StorageService {
//	return &StorageService{
//		VideoDal: &implement.VideoDal{},
//		UserDal:  &implement.UserDal{},
//	}
//}

func (v *StorageService) UpdateAvatar(uid uint, url string) error {
	return v.UserDal.UpdateAvatar(uid, url)
}

func (v *StorageService) UploadVideoCallback(authorId uint, url string, coverUrl string, describe string, videoTypeId uint) (err error) {
	video := &model.Video{
		AuthorID:    authorId,
		PlayUrl:     url,
		CoverUrl:    coverUrl,
		VideoTypeID: videoTypeId,
		Describe:    describe,
	}
	_, err = v.VideoDal.CreateVideo(video)
	if err != nil {
		return err
	}
	return
}

// 实现接口
var _ abstract2.StorageService = (*StorageService)(nil)
