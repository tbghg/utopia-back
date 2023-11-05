package abstract

type StorageService interface {
	// UploadVideoCallback 上传文件七牛云回调
	UploadVideoCallback(uid uint, url string, coverUrl string, describe string, title string, videoTypeId uint) error
	// UpdateAvatar 更新头像
	UpdateAvatar(uid uint, url string) error
}
