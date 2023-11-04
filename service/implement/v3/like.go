package v3

import (
	"fmt"
	"utopia-back/cache"
	"utopia-back/database/abstract"
	"utopia-back/pkg/logger"
	abstract2 "utopia-back/service/abstract"
)

type LikeService struct {
	LikeDal  abstract.LikeDal
	VideoDal abstract.VideoDal
}

func (l LikeService) Like(userId uint, videoId uint) (err error) {
	//判断videoId是否存在
	err = l.VideoDal.IsVideoExist(videoId)
	if err != nil {
		return err
	}

	status := cache.HSetUserLikedVideo(userId, []uint{videoId})

	//更新数据库
	rowsAffected, err := l.LikeDal.Like(userId, videoId)
	if err != nil {
		return err
	}

	if status == 1 {
		go l.rebuildUserLikedVideoCache(userId)
	}

	if rowsAffected != 0 {
		// 更新缓存
		key := cache.VideoLikeCountKey(videoId)
		err = cache.RDB.Incr(cache.Ctx, key).Err()
		if err != nil {
			return err
		}
		l.LikeDal.UpdateLikeCount(videoId, 1)
	}

	return nil
}

func (l LikeService) UnLike(userId uint, videoId uint) (err error) {
	//判断videoId是否存在
	err = l.VideoDal.IsVideoExist(videoId)
	if err != nil {
		return err
	}

	cache.HDelUserLikedVideo(userId, []uint{videoId})

	//更新数据库
	rowsAffected, err := l.LikeDal.UnLike(userId, videoId)
	if err != nil {
		return err
	}

	if rowsAffected != 0 {
		// 更新缓存
		key := cache.VideoLikeCountKey(videoId)
		err = cache.RDB.Decr(cache.Ctx, key).Err()
		if err != nil {
			return err
		}
		l.LikeDal.UpdateLikeCount(videoId, -1)
	}
	return
}

// 构建缓存
func (l LikeService) rebuildUserLikedVideoCache(userId uint) {
	videoIds, err := l.LikeDal.GetUserLikedVideosWithLimit(userId, cache.SetFieldNum)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("IsLike l.LikeDal.GetUserLikedVideosWithLimit uid:%v err:%+v", userId, err))
		return
	}
	minVid := uint(0)
	if len(videoIds) > cache.SetFieldNum {
		minVid = videoIds[cache.SetFieldNum-1]
		videoIds = videoIds[:cache.SetFieldNum]
	}
	cache.BuildUserLikedVideos(cache.UserLikedVideoKeyV3(userId), videoIds, minVid)
}

var _ abstract2.LikeService = (*LikeService)(nil)
