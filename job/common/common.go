package common

import "sync/atomic"

// PopularVideoVersion 当前缓存版本
var PopularVideoVersion atomic.Int64

// GetPopularVideoVersion 获取当前热门视频版本
func GetPopularVideoVersion() int {
	return int(PopularVideoVersion.Load())
}
