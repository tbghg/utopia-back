package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"utopia-back/http/controller/base"
	"utopia-back/model"
	"utopia-back/pkg/logger"
	"utopia-back/service/abstract"
)

type VideoController struct {
	VideoService abstract.VideoService
}

type VideoResp struct {
	VideoInfo []*model.VideoInfo `json:"video_info"`
	NextTime  int                `json:"next_time"`
}

type SingleVideo struct {
	VideoInfo *model.VideoInfo `json:"video_info"`
}

type VideoRespWithoutTime struct {
	VideoInfo []*model.VideoInfo `json:"video_info"`
}

type PopularVideoResp struct {
	VideoInfo []*model.VideoInfo `json:"video_info"`
	Score     float64            `json:"score"`
	Version   int                `json:"version"`
}

func (v VideoController) GetCategoryVideos(c *gin.Context) {
	var (
		err        error
		videoInfos []*model.VideoInfo
		nextTime   int
		uid        int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("GetCategoryVideos err:%+v", err))
		}
	}()

	sLastTime, _ := c.GetQuery("last_time")
	lastTime, err := strconv.ParseInt(sLastTime, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("last_time:%v 参数错误", sLastTime))
		return
	}
	sVideoTypeId, _ := c.GetQuery("video_type_id")
	videoTypeId, err := strconv.ParseInt(sVideoTypeId, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("video_type_id:%v 参数错误", sVideoTypeId))
		return
	}
	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfos, nextTime, err = v.VideoService.GetCategoryVideos(uint(uid), uint(lastTime), uint(videoTypeId))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: VideoResp{
			VideoInfo: videoInfos,
			NextTime:  nextTime,
		},
	})
}

// GetPopularVideos 获取热门视频
func (v VideoController) GetPopularVideos(c *gin.Context) {
	var (
		err              error
		videoInfos       []*model.VideoInfo
		uid              int
		score, nextScore float64
		nextVersion      int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("GetUploadVideos err:%+v", err))
		}
	}()

	sVersion, _ := c.GetQuery("version")
	version, err := strconv.ParseInt(sVersion, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("version:%v 参数错误", sVersion))
		return
	}

	sScore, _ := c.GetQuery("score")
	score, err = strconv.ParseFloat(sScore, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("score:%v 参数错误", sScore))
		return
	}

	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfos, nextScore, nextVersion, err = v.VideoService.GetPopularVideos(uint(uid), int(version), score)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: PopularVideoResp{
			VideoInfo: videoInfos,
			Score:     nextScore,
			Version:   nextVersion,
		},
	})
}

// GetRecommendVideos 获取推荐视频
func (v VideoController) GetRecommendVideos(c *gin.Context) {
	var (
		err        error
		videoInfos []*model.VideoInfo
		nextTime   int
		uid        int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("GetRecommendVideos err:%+v", err))
		}
	}()

	sLastTime, _ := c.GetQuery("last_time")
	lastTime, err := strconv.ParseInt(sLastTime, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("last_time:%v 参数错误", sLastTime))
		return
	}
	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfos, nextTime, err = v.VideoService.GetRecommendVideos(uint(uid), uint(lastTime))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: VideoResp{
			VideoInfo: videoInfos,
			NextTime:  nextTime,
		},
	})
	return
}

// GetUploadVideos 获取上传的视频
func (v VideoController) GetUploadVideos(c *gin.Context) {
	var (
		err        error
		videoInfos []*model.VideoInfo
		nextTime   int
		uid        int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("GetUploadVideos err:%+v", err))
		}
	}()

	sLastTime, _ := c.GetQuery("last_time")
	lastTime, err := strconv.ParseInt(sLastTime, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("last_time:%v 参数错误", sLastTime))
		return
	}
	sTargetUserId, _ := c.GetQuery("user_id")
	targetUid, err := strconv.ParseInt(sTargetUserId, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("user_id:%v 参数错误", sTargetUserId))
		return
	}
	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfos, nextTime, err = v.VideoService.GetUploadVideos(uint(uid), uint(targetUid), uint(lastTime))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: VideoResp{
			VideoInfo: videoInfos,
			NextTime:  nextTime,
		},
	})
}

// GetSingleVideo 获取上传的视频
func (v VideoController) GetSingleVideo(c *gin.Context) {
	var (
		err       error
		videoInfo *model.VideoInfo
		uid       int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("GetSingleVideos err:%+v", err))
		}
	}()

	sVideoId, _ := c.GetQuery("video_id")
	vid, err := strconv.ParseInt(sVideoId, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("video_id:%v 参数错误", vid))
		return
	}
	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfo, err = v.VideoService.GetSingleVideo(uint(uid), uint(vid))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: SingleVideo{
			VideoInfo: videoInfo,
		},
	})
}

// GetFavoriteVideos 获取收藏的视频
func (v VideoController) GetFavoriteVideos(c *gin.Context) {
	var (
		err        error
		videoInfos []*model.VideoInfo
		nextTime   int
		uid        int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("GetFavoriteVideos err:%+v", err))
		}
	}()

	sLastTime, _ := c.GetQuery("last_time")
	lastTime, err := strconv.ParseInt(sLastTime, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("last_time:%v 参数错误", sLastTime))
		return
	}
	sTargetUserId, _ := c.GetQuery("user_id")
	targetUid, err := strconv.ParseInt(sTargetUserId, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("user_id:%v 参数错误", sLastTime))
		return
	}
	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfos, nextTime, err = v.VideoService.GetFavoriteVideos(uint(uid), uint(targetUid), uint(lastTime))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: VideoResp{
			VideoInfo: videoInfos,
			NextTime:  nextTime,
		},
	})
}

// SearchVideos 搜索视频
func (v VideoController) SearchVideos(c *gin.Context) {
	var (
		err        error
		videoInfos []*model.VideoInfo
		uid        int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("SearchVideos err:%+v", err))
		}
	}()

	search, ok := c.GetQuery("search")
	if !ok {
		err = errors.New("search 参数错误")
		return
	}

	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfos, err = v.VideoService.SearchVideos(uint(uid), search)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: VideoRespWithoutTime{
			VideoInfo: videoInfos,
		},
	})
}
