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

type FollowController struct {
	FollowService abstract.FollowService
}

type followRequest struct {
	ToUserId   int `json:"to_user_id" validate:"required,gt=0"` // 要关注的用户id > 0 必需
	ActionType int `json:"action_type" validate:"required"`     // 操作类型 1: 关注 2: 取消关注
}

// Follow 关注/取消关注
func (f *FollowController) Follow(c *gin.Context) {
	var (
		r   followRequest
		err error
	)

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("Follow err:%+v", err))
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
		return
	}

	if userId == r.ToUserId {
		c.JSON(http.StatusOK, base.SuccessResponse)
		return
	}
	// 判断操作类型
	switch r.ActionType {
	case 1:
		err = f.FollowService.Follow(uint(userId), uint(r.ToUserId))
	case 2:
		err = f.FollowService.UnFollow(uint(userId), uint(r.ToUserId))
	default:
		err = base.ActionTypeInvalidError
	}

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, base.SuccessResponse)

}

// FansList 获取粉丝列表
func (f *FollowController) FansList(c *gin.Context) {
	var err error

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			// 记录错误日志
			logger.Logger.Error(fmt.Sprintf("FansList err:%+v", err))
		}
	}()

	// 获取用户id
	value, ok := c.Get("user_id")
	userId, ok := value.(int)
	if !ok {
		err = base.UserIDInvalidError
		return
	}

	list, err := f.FollowService.GetFansList(uint(userId))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &base.FollowResponse{
		Code:     base.SuccessCode,
		Msg:      "ok",
		UserList: list,
	})
}

// FollowList 获取关注列表
func (f *FollowController) FollowList(c *gin.Context) {
	var err error

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			// 记录错误日志
			logger.Logger.Error(fmt.Sprintf("FollowList err:%+v", err))
		}
	}()

	// 获取用户id
	value, ok := c.Get("user_id")
	userId, ok := value.(int)
	if !ok {
		err = base.UserIDInvalidError
		return
	}

	list, err := f.FollowService.GetFollowList(uint(userId))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &base.FollowResponse{
		Code:     base.SuccessCode,
		Msg:      "ok",
		UserList: list,
	})
}
