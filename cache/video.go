package cache

import "fmt"

// PopularVideoKey 热门视频Key
func PopularVideoKey() string {
	return "video:popular"
}

// VideoInfoKey 视频信息Key
func VideoInfoKey(videoId uint) string {
	return fmt.Sprintf("video:info:%v", videoId)
}
