package job

import (
	"utopia-back/database/implement"
	v1 "utopia-back/service/implement/v1"
)

type videoJob struct {
	videoServer *v1.VideoService
}

func videoJobInit(dal *implement.CenterDal) {
	v := &videoJob{
		videoServer: &v1.VideoService{
			VideoDal:    dal.VideoDal,
			UserDal:     dal.UserDal,
			FollowDal:   dal.FollowDal,
			LikeDal:     dal.LikeDal,
			FavoriteDal: dal.FavoriteDal,
		},
	}

	go v.updatePopularVideos()
}

// 热门视频缓存 每隔一小时更新一次
// 存入zset，通过score判断此时取到的位置
// 需要一个version来标记缓存版本是否更新，如更新忽略score直接重新取
func (v *videoJob) updatePopularVideos() {
}
