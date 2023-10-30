package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"utopia-back/http/controller/base"
	"utopia-back/model"
	"utopia-back/service/abstract"
)

type VideoController struct {
	VideoService abstract.VideoService
}

type VideoResp struct {
	VideoInfo []*model.VideoInfo `json:"Video_info"`
	NextTime  string             `json:"next_time"`
}

func (v VideoController) GetCategoryVideos(c *gin.Context) {
	var (
		err        error
		videoInfos []*model.VideoInfo
		nextTime   string
		uid        int
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
		}
	}()

	lastTime, _ := c.GetQuery("last_time")
	sVideoTypeId, _ := c.GetQuery("video_type_id")
	videoTypeId, err := strconv.ParseInt(sVideoTypeId, 10, 64)
	if err != nil {
		err = errors.New("参数错误")
		return
	}
	// 获取用户id
	if value, ok := c.Get("user_id"); ok {
		uid, _ = value.(int)
	}

	videoInfos, nextTime, err = v.VideoService.GetCategoryVideos(uint(uid), lastTime, uint(videoTypeId))
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
