package v1

import (
	"utopia-back/database/abstract"
	"utopia-back/model"
	abstract2 "utopia-back/service/abstract"
)

type VideoService struct {
	VideoDal    abstract.VideoDal
	UserDal     abstract.UserDal
	FollowDal   abstract.FollowDal
	FavoriteDal abstract.FavoriteDal
	LikeDal     abstract.LikeDal
}

func (v VideoService) GetCategoryVideos(uid uint, lastTime uint, videoTypeId uint) ([]*model.VideoInfo, int, error) {
	var videoInfos []*model.VideoInfo
	videos, err := v.VideoDal.GetVideoByType(lastTime, videoTypeId)
	if err != nil || len(videos) == 0 {
		return nil, -1, err
	}

	videoInfos = make([]*model.VideoInfo, len(videos))
	for i := range videos {
		videoInfos[i] = new(model.VideoInfo)
		videoInfos[i].ID = videos[i].ID
		videoInfos[i].CreatedAt = videos[i].CreatedAt
		videoInfos[i].PlayUrl = videos[i].PlayUrl
		videoInfos[i].CoverUrl = videos[i].CoverUrl
		videoInfos[i].VideoTypeID = videos[i].VideoTypeID
		videoInfos[i].Describe = videos[i].Describe

		userInfo, err := v.UserDal.GetUserInfoById(videos[i].AuthorID)
		if err != nil {
			return nil, -1, err
		}

		videoInfos[i].Author.ID = userInfo.ID
		videoInfos[i].Author.Nickname = userInfo.Nickname
		videoInfos[i].Author.Avatar = userInfo.Avatar
		videoInfos[i].Author.Username = userInfo.Username
		videoInfos[i].Author.FansCount = userInfo.FansCount
		videoInfos[i].Author.FollowCount = userInfo.FollowCount
		videoInfos[i].Author.VideoCount = userInfo.VideoCount

		if uid != 0 {
			videoInfos[i].IsFollow, err = v.FollowDal.IsFollow(uid, videos[i].AuthorID)
			if err != nil {
				return nil, -1, err
			}
			videoInfos[i].IsLike, err = v.LikeDal.IsLike(uid, videos[i].AuthorID)
			if err != nil {
				return nil, -1, err
			}
			videoInfos[i].IsFavorite, err = v.FavoriteDal.IsFavorite(uid, videos[i].AuthorID)
			if err != nil {
				return nil, -1, err
			}
		}
	}
	nextTime := videoInfos[len(videoInfos)-1].CreatedAt.UnixMilli()
	return videoInfos, int(nextTime), nil
}

var _ abstract2.VideoService = (*VideoService)(nil)
