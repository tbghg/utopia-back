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

func (v *videoJob) updatePopularVideos() {

}
