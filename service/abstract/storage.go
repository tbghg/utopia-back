package abstract

import "utopia-back/model"

type StorageService interface {
	// UploadVideoCallback 上传文件七牛云回调
	UploadVideoCallback(uid uint, url string, coverUrl string, describe string, title string, videoTypeId uint, isWithCover bool, key string) error
	// UpdateAvatar 更新头像
	UpdateAvatar(uid uint, url string) error
	// PreVideoCallback 视频预处理回调
	PreVideoCallback(inputKey string, item []model.CallbackItem) error
}
