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
	utils "utopia-back/pkg/util"
	"utopia-back/service/abstract"
)

type CommentController struct {
	CommentService abstract.CommentService
}

type CommentReq struct {
	VideoId uint   `json:"video_id"`
	Content string `json:"content"`
}

type CommentListResp struct {
	CommentInfo []*model.CommentInfo `json:"comment_info"`
	NextTime    int                  `json:"next_time"`
}

// Comment 点赞/取消点赞
func (l *CommentController) Comment(c *gin.Context) {
	var (
		err error
		r   CommentReq
	)

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("Comment err:%+v", err))
		}
	}()

	// 获取用户id
	value, ok := c.Get("user_id")
	userId, ok := value.(int)
	if !ok {
		err = base.UserIDInvalidError
		return
	}

	// 接收参数并绑定
	if err = c.ShouldBindJSON(&r); err != nil {
		return
	}
	// 参数校验
	if err = utils.Validate.Struct(r); err != nil {
		err = base.ValidParameterError
		return
	}
	err = l.CommentService.Comment(uint(userId), r.VideoId, r.Content)
	if err != nil {
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, base.SuccessResponse)
}

// CommentList 评论列表
func (l *CommentController) CommentList(c *gin.Context) {
	var err error

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("CommentList err:%+v", err))
		}
	}()

	sLastTime, _ := c.GetQuery("last_time")
	lastTime, err := strconv.ParseInt(sLastTime, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("last_time:%v 参数错误", sLastTime))
		return
	}

	// 参数校验
	sVideoId, _ := c.GetQuery("video_id")
	videoId, err := strconv.ParseInt(sVideoId, 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("video_id:%v 参数错误", videoId))
		return
	}

	list, nextTime, err := l.CommentService.CommentList(uint(videoId), uint(lastTime))
	if err != nil {
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: CommentListResp{
			list,
			nextTime,
		},
	})
}
