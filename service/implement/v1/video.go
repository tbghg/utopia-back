package v1

import (
	"fmt"
	"sync"
	"utopia-back/cache"
	"utopia-back/database/abstract"
	"utopia-back/model"
	"utopia-back/pkg/logger"
	abstract2 "utopia-back/service/abstract"
)

type VideoService struct {
	VideoDal    abstract.VideoDal
	UserDal     abstract.UserDal
	FollowDal   abstract.FollowDal
	FavoriteDal abstract.FavoriteDal
	LikeDal     abstract.LikeDal
}

const (
	// -1为不限制
	// todo 为方便调试，先写为3，之后改为20
	favoriteVideosLimit  = 3
	uploadVideosLimit    = 3
	typeVideosLimit      = 3
	popularVideosLimit   = 3
	recommendVideosLimit = 3
)

func (v VideoService) GetPopularVideos(uid uint, lastTime uint) ([]*model.VideoInfo, int, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoService) GetRecommendVideos(uid uint, lastTime uint) ([]*model.VideoInfo, int, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoService) GetFavoriteVideos(uid uint, targetUid uint, lastTime uint) (videoInfos []*model.VideoInfo, nextTime int, err error) {
	var videoIds []uint
	videoIds, nextTime, err = v.FavoriteDal.GetFavoriteList(targetUid, lastTime, favoriteVideosLimit)
	if err != nil || len(videoIds) == 0 {
		return nil, -1, err
	}
	videos, err := v.VideoDal.GetVideoInfoById(videoIds)
	if err != nil {
		return nil, 0, err
	}
	videoInfos, _, err = v.getVideoInfo(uid, videos)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (v VideoService) GetUploadVideos(uid uint, targetUid uint, lastTime uint) ([]*model.VideoInfo, int, error) {
	videos, err := v.VideoDal.GetUploadVideos(lastTime, targetUid, uploadVideosLimit)
	if err != nil || len(videos) == 0 {
		return nil, -1, err
	}
	return v.getVideoInfo(uid, videos)
}

func (v VideoService) SearchVideoAndUser(search string) ([]*model.VideoInfo, []*model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (v VideoService) GetCategoryVideos(uid uint, lastTime uint, videoTypeId uint) ([]*model.VideoInfo, int, error) {
	videos, err := v.VideoDal.GetVideoByType(lastTime, videoTypeId, typeVideosLimit)
	if err != nil || len(videos) == 0 {
		return nil, -1, err
	}
	return v.getVideoInfo(uid, videos)
}

func (v VideoService) getVideoInfo(userId uint, videos []*model.Video) ([]*model.VideoInfo, int, error) {
	var (
		wg         sync.WaitGroup
		liked      map[uint]bool // 用户是否为该视频点赞
		videoInfos = make([]*model.VideoInfo, len(videos))
		//authorInfos = make(map[uint]*model.UserInfo)
		authorInfos sync.Map
		videoIds    = make([]uint, 0, len(videos))
		authorIds   = make([]uint, 0)
	)

	for i, video := range videos {
		videoInfos[i] = new(model.VideoInfo)
		videoIds = append(videoIds, video.ID)
		_, ok := authorInfos.LoadOrStore(video.AuthorID, nil)
		if !ok {
			authorIds = append(authorIds, video.AuthorID)
		}
	}

	// 判断是否为该视频点过赞
	if userId != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			liked, err = v.batchIsUserLiked(userId, videoIds)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("getVideoInfo v.batchIsUserLiked userId:%v videoIds:%v err:%+v",
					userId, videoIds, err))
			}
			return
		}()
	}

	// 获取点赞数/收藏数 是否收藏
	for i := range videos {
		wg.Add(1)
		vid := videos[i].ID
		i := i
		go func() {
			defer wg.Done()
			res, err := cache.GetStringCache(func() (interface{}, error) {
				return v.LikeDal.GetLikeCount(vid)
			}, cache.VideoLikeCountKey(vid), cache.TypeInt64)

			if err != nil {
				logger.Logger.Error(fmt.Sprintf("getVideoInfo cache.GetStringCache userId:%v vid:%v err:%+v",
					userId, vid, err))
			}

			videoInfos[i].LikeCount = int(res.(int64))

			favoriteCount, err := v.FavoriteDal.GetFavoriteCount(vid)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("getVideoInfo v.FavoriteDal.GetFavoriteCount vid:%v err:%+v",
					vid, err))
			}
			videoInfos[i].FavoriteCount = int(favoriteCount)

			if userId != 0 {
				// 是否收藏
				videoInfos[i].IsFavorite, err = v.FavoriteDal.IsFavorite(userId, videos[i].ID)
				if err != nil {
					logger.Logger.Error(fmt.Sprintf("v.FavoriteDal.IsFavorite userId:%v authorId:%v err:%+v", userId, videos[i].AuthorID, err))
				}
			}
		}()

	}

	// 查询视频作者信息, 是否关注过该用户
	for _, uid := range authorIds {
		uid := uid
		wg.Add(1)
		go func() {
			defer wg.Done()
			userInfo, err := v.UserDal.GetUserInfoById(uid)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("getVideoInfo v.UserDal.GetUserInfoById userId:%v err:%+v",
					userId, err))
			}

			var isFollow bool
			if userId != 0 {
				isFollow, err = v.FollowDal.IsFollow(userId, uid)
				if err != nil {
					logger.Logger.Error(fmt.Sprintf("getVideoInfo v.FollowDal.IsFollow userId:%v authorId:%v err:%+v",
						userId, uid, err))
				}
			}

			authorInfos.Store(uid, &struct {
				userInfo *model.UserInfo
				isFollow bool
			}{userInfo: &userInfo, isFollow: isFollow})
		}()
	}

	wg.Wait()

	for i := range videos {
		// 视频信息
		videoInfos[i].ID = videos[i].ID
		videoInfos[i].CreatedAt = videos[i].CreatedAt
		videoInfos[i].PlayUrl = videos[i].PlayUrl
		videoInfos[i].CoverUrl = videos[i].CoverUrl
		videoInfos[i].VideoTypeID = videos[i].VideoTypeID
		videoInfos[i].Describe = videos[i].Describe

		// 作者信息 是否关注作者
		if userInfoStructInterface, ok := authorInfos.Load(videos[i].AuthorID); ok {
			userInfoStruct := userInfoStructInterface.(*struct {
				userInfo *model.UserInfo
				isFollow bool
			})
			userInfo := userInfoStruct.userInfo

			videoInfos[i].Author.ID = userInfo.ID
			videoInfos[i].Author.Nickname = userInfo.Nickname
			videoInfos[i].Author.Avatar = userInfo.Avatar
			videoInfos[i].Author.Username = userInfo.Username
			videoInfos[i].Author.FansCount = userInfo.FansCount
			videoInfos[i].Author.FollowCount = userInfo.FollowCount
			videoInfos[i].Author.VideoCount = userInfo.VideoCount

			videoInfos[i].IsFollow = userInfoStruct.isFollow
		}

		// 是否点赞
		if lk, ok := liked[videos[i].ID]; ok {
			videoInfos[i].IsLike = lk
		}
	}
	nextTime := videoInfos[len(videoInfos)-1].CreatedAt.UnixMilli()
	return videoInfos, int(nextTime), nil
}

// 用户是否给视频点过赞
func (v VideoService) batchIsUserLiked(userId uint, videoIds []uint) (isLikeMap map[uint]bool, err error) {
	// 查询缓存
	result, state := cache.IsUserLikedVideos(userId, videoIds)

	resourceDB := func(userId uint, videoIds []uint) (isLikeMap map[uint]bool, err error) {
		var likedVideoIds []uint
		// 不存在该Key 查DB
		likedVideoIds, err = v.LikeDal.BatchIsLike(userId, videoIds)
		// 查询DB失败，直接返回报错
		if err != nil {
			return nil, err
		}
		isLikeMap = make(map[uint]bool, len(likedVideoIds))
		for _, vid := range likedVideoIds {
			isLikeMap[vid] = true
		}
		return
	}

	switch state {
	case 1:
		// 不存在该key，需重新创建
		// 开启协程 异步构建
		go v.rebuildUserLikedVideoCache(userId)
		return resourceDB(userId, videoIds)
	case 2:
		// cache异常，回源不写cache
		return resourceDB(userId, videoIds)
	case 0:
		// 查询成功
		// 0未点赞；1为点赞；2 冷数据,需回源
		isLikeMap = make(map[uint]bool, len(result))
		codeVideoIds := make([]uint, 0)
		for v, s := range result {
			if s == 2 { // 冷数据单独处理
				codeVideoIds = append(codeVideoIds, v)
			} else {
				isLikeMap[v] = s == 1
			}
		}
		// 冷数据回源
		likedVideoIds, err := v.LikeDal.BatchIsLike(userId, codeVideoIds)
		if err != nil {
			return nil, err
		}
		for _, vid := range likedVideoIds {
			isLikeMap[vid] = true
		}
	}
	return
}

// 构建缓存
func (v VideoService) rebuildUserLikedVideoCache(userId uint) {
	videoIds, err := v.LikeDal.GetUserLikedVideosWithLimit(userId, cache.SetFieldNum)
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

var _ abstract2.VideoService = (*VideoService)(nil)
