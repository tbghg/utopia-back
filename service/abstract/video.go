package abstract

type VideoService interface {
	// UploadVideoCallback 上传文件七牛云回调
	UploadVideoCallback(authorId uint, url string, coverUrl string, describe string, videoType uint) error
	// UpdateAvatar 更新头像
	UpdateAvatar(uid uint, url string) error
}
