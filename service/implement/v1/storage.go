package v1

import (
	"fmt"
	"utopia-back/cache"
	"utopia-back/config"
	"utopia-back/database/abstract"
	"utopia-back/model"
	"utopia-back/pkg/logger"
	abstract2 "utopia-back/service/abstract"
)

type StorageService struct {
	VideoDal abstract.VideoDal
	UserDal  abstract.UserDal
}

func (v *StorageService) PreVideoCallback(inputKey string, item []model.CallbackItem) error {
	vid := cache.GetVideoPersistent(inputKey)
	for _, value := range item {
		if value.Code != 0 {
			logger.Logger.Error(fmt.Sprintf("PreVideoCallback item:%+v", item))
			break
		}
		coverUrl := config.V.GetString("qiniu.kodoApi") + "/" + value.Key
		err := v.VideoDal.UpdateCover(vid, coverUrl)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("PreVideoCallback v.VideoDal.UpdateCover err:&+v", err))
		}
	}
	return nil
}

func (v *StorageService) UpdateAvatar(uid uint, url string) error {
	return v.UserDal.UpdateAvatar(uid, url)
}

func (v *StorageService) UploadVideoCallback(authorId uint, url string, coverUrl string, describe string, title string, videoTypeId uint, isWithCover bool, key string) (err error) {
	video := &model.Video{
		AuthorID:    authorId,
		PlayUrl:     url,
		CoverUrl:    coverUrl,
		VideoTypeID: videoTypeId,
		Describe:    describe,
		Title:       title,
	}
	vid, err := v.VideoDal.CreateVideo(video)
	if err != nil {
		return err
	}
	// 不携带封面则写入缓存
	if !isWithCover {
		cache.SetVideoPersistent(key, vid)
	}
	return
}

// 实现接口
var _ abstract2.StorageService = (*StorageService)(nil)
