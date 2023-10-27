package abstract

type VideoService interface {
	// UploadVideoCallback 上传视频回调
	UploadVideoCallback(authorId uint, url string, coverUrl string, describe string) error
}
