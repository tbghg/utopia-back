package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"utopia-back/http/controller/base"
	"utopia-back/pkg/logger"
	utils "utopia-back/pkg/util"
	"utopia-back/service/abstract"
)

type LikeController struct {
	LikeService abstract.LikeService
}

type likeRequest struct {
	VideoId    uint `json:"video_id" validate:"required,gt=0"` // 视频id > 0 必需
	ActionType int  `json:"action_type" validate:"required"`   // 操作类型 1: 点赞 2: 取消点赞
}

// Like 点赞/取消点赞
func (l *LikeController) Like(c *gin.Context) {
	var (
		r   likeRequest
		err error
	)

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("Like err:%+v", err))
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
	// 判断操作类型
	switch r.ActionType {
	case 1:
		err = l.LikeService.Like(uint(userId), r.VideoId)
	case 2:
		err = l.LikeService.UnLike(uint(userId), r.VideoId)
	default:
		err = base.ActionTypeInvalidError
	}
	if err != nil {
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, base.SuccessResponse)
}
