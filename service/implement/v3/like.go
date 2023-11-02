package v3

//import (
//	"fmt"
//	"utopia-back/cache"
//	"utopia-back/database/abstract"
//	"utopia-back/pkg/logger"
//	abstract2 "utopia-back/service/abstract"
//)
//
//type LikeService struct {
//	LikeDal  abstract.LikeDal
//	VideoDal abstract.VideoDal
//}
//
//func (l LikeService) Like(userId uint, videoId uint) (err error) {
//	//判断videoId是否存在
//	err = l.VideoDal.IsVideoExist(videoId)
//	if err != nil {
//		return err
//	}
//	//更新数据库
//	err = l.LikeDal.Like(userId, videoId)
//	if err != nil {
//		return err
//	}
//	// 更新缓存
//	key := cache.VideoLikeCountKey(videoId)
//	err = cache.RDB.Incr(cache.Ctx, key).Err()
//
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (l LikeService) UnLike(userId uint, videoId uint) (err error) {
//	//判断videoId是否存在
//	err = l.VideoDal.IsVideoExist(videoId)
//	if err != nil {
//		return err
//	}
//	//更新数据库
//	err = l.LikeDal.UnLike(userId, videoId)
//	if err != nil {
//		return err
//	}
//	// 更新缓存
//	key := cache.VideoLikeCountKey(videoId)
//	err = cache.RDB.Decr(cache.Ctx, key).Err()
//	if err != nil {
//		return err
//	}
//	return l.LikeDal.UnLike(userId, videoId)
//}
//
//func (l LikeService) IsLike(userId uint, videoId uint) (isLike bool, err error) {
//	// 查询缓存
//	liked, state := cache.IsUserLikedVideo(userId, videoId)
//	switch state {
//	case 0:
//		return liked, nil
//	case 1:
//		// 不存在该key，需构建
//		isLike, err = l.LikeDal.IsLike(userId, videoId)
//		go l.rebuildUserLikedVideoCache(userId)
//		return
//	case 2:
//		// 冷数据，回源+写回cache
//		isLike, err = l.LikeDal.IsLike(userId, videoId)
//		cache.HSetUserLikedVideo(userId, []uint{videoId})
//		return
//	case 3:
//		// redis异常，回源不写回
//		return l.LikeDal.IsLike(userId, videoId)
//	}
//	return
//}
//
//// BatchIsLike 批量查询点赞接口
//func (l LikeService) BatchIsLike(userId uint, videoIds []uint) (isLikeMap map[uint]bool, err error) {
//	// 查询缓存
//	result, state := cache.IsUserLikedVideos(userId, videoIds)
//	switch state {
//	case 1:
//		// 不存在该Key，重新构建
//		likedVideoIds, err := l.LikeDal.BatchIsLike(userId, videoIds)
//		if err != nil {
//			return nil, err
//		}
//		isLikeMap = make(map[uint]bool, len(likedVideoIds))
//		for _, vid := range likedVideoIds {
//			isLikeMap[vid] = true
//		}
//		go l.rebuildUserLikedVideoCache(userId)
//		return
//	case 2:
//		// cache异常，回源不写cache
//		// 不存在该Key，重新构建
//		likedVideoIds, err := l.LikeDal.BatchIsLike(userId, videoIds)
//		if err != nil {
//			return nil, err
//		}
//		isLikeMap = make(map[uint]bool, len(likedVideoIds))
//		for _, vid := range likedVideoIds {
//			isLikeMap[vid] = true
//		}
//		return
//	case 0:
//		// 查询成功
//		// 0未点赞；1为点赞；2 冷数据,需回源
//		isLikeMap = make(map[uint]bool, len(result))
//		codeVideoIds := make([]uint, 0)
//		for v, s := range result {
//			if s == 2 { // 冷数据单独处理
//				codeVideoIds = append(codeVideoIds, v)
//			} else {
//				isLikeMap[v] = s == 1
//			}
//		}
//		// 冷数据回源
//		likedVideoIds, err := l.LikeDal.BatchIsLike(userId, codeVideoIds)
//		if err != nil {
//			return nil, err
//		}
//		for _, vid := range likedVideoIds {
//			isLikeMap[vid] = true
//		}
//	}
//	return
//}
//
//func (l LikeService) rebuildUserLikedVideoCache(userId uint) {
//	videoIds, err := l.LikeDal.GetUserLikedVideosWithLimit(userId, cache.SetFieldNum)
//	if err != nil {
//		logger.Logger.Error(fmt.Sprintf("IsLike l.LikeDal.GetUserLikedVideosWithLimit uid:%v err:%+v", userId, err))
//		return
//	}
//	minVid := uint(0)
//	if len(videoIds) > cache.SetFieldNum {
//		minVid = videoIds[cache.SetFieldNum-1]
//	}
//	cache.BuildUserLikedVideos(cache.UserLikedVideoKeyV3(userId), videoIds[:cache.SetFieldNum], minVid)
//}
//
//func (l LikeService) GetLikeCount(videoId uint) (count int64, err error) {
//	// 构造key
//	key := cache.VideoLikeCountKey(videoId)
//	// 缓存层取数据
//	res, err := cache.GetStringCache(func() (interface{}, error) { return l.LikeDal.GetLikeCount(videoId) }, key, cache.TypeInt64)
//	// 返回结果
//	if err != nil {
//		return 0, err
//	}
//	return res.(int64), nil
//}
//
//var _ abstract2.LikeService = (*LikeService)(nil)
