package job

import (
	"fmt"
	"strconv"
	"time"
	"utopia-back/cache"
	"utopia-back/database/implement"
	"utopia-back/job/common"
	"utopia-back/pkg/logger"
	v1 "utopia-back/service/implement/v1"
)

type videoJob struct {
	videoServer *v1.VideoService
}

const popularVideosNum = 100

func VideoJobInit(dal *implement.CenterDal) {
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

// 热门视频缓存 每隔三小时更新一次
// 存入zset，通过score判断此时取到的位置
// 需要一个version来标记缓存版本是否更新，如更新忽略score直接重新取
func (v *videoJob) updatePopularVideos() {
	var nums int64
	ticker := time.NewTicker(3 * time.Hour)
	for {
		nums = (nums + 1) % 2 // count值为0/1
		common.PopularVideoVersion.Store(nums + 1)

		videoIds, err := v.videoServer.VideoDal.GetPopularVideos(popularVideosNum)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("updatePopularVideos v.videoServer.VideoDal.GetPopularVideos err:%+v", videoIds))
			// 一分钟后重试
			<-time.After(5 * time.Minute)
			continue
		}
		popularItem := make([]*cache.VideoPopularItem, len(videoIds))
		for i := range videoIds {
			sScore := fmt.Sprintf("%d.%d", videoIds[i].Count, videoIds[i].VideoID)

			popularItem[i] = &cache.VideoPopularItem{Vid: videoIds[i].VideoID}
			popularItem[i].Score, _ = strconv.ParseFloat(sScore, 64)
		}

		key := cache.PopularVideoKey(common.GetPopularVideoVersion())
		cache.BuildPopularVideo(key, popularItem)

		<-ticker.C
	}
}
