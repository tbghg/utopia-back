package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	utils "utopia-back/pkg/util"
	"utopia-back/service/abstract"
)

type FollowController struct {
	Service abstract.FollowService
}

func NewFollowController(service abstract.FollowService) *FollowController {
	return &FollowController{Service: service}
}

type followRequest struct {
	toUserId   int `form:"to_user_id" validate:"required,gt=0"` // 要关注的用户id > 0 必需
	actionType int `form:"action_type" validate:"required"`     // 操作类型 1: 关注 2: 取消关注
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
			c.JSON(http.StatusOK, &ResponseWithoutData{
				Code: ErrorCode,
				Msg:  err.Error(),
			})
		}
	}()

	// 获取用户id
	value, ok := c.Get("user_id")
	userId, ok := value.(int)
	if !ok {
		err = UserIDInvalidError
		return
	}

	// 接收参数并绑定
	if err = c.ShouldBindQuery(&r); err != nil {
		return
	}
	// 参数校验
	if err = utils.Validate.Struct(r); err != nil {
		return
	}
	// 判断操作类型
	switch r.actionType {
	case 1:
		err = f.Service.Follow(uint(userId), uint(r.toUserId))
	case 2:
		err = f.Service.UnFollow(uint(userId), uint(r.toUserId))
	default:
		err = ActionTypeInvalidError
	}

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, SuccessResponse)

}

// FansList 获取粉丝列表
func (f *FollowController) FansList(c *gin.Context) {
	var err error

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &ResponseWithoutData{
				Code: ErrorCode,
				Msg:  err.Error(),
			})
		}
	}()

	// 获取用户id
	value, ok := c.Get("user_id")
	userId, ok := value.(int)
	if !ok {
		err = UserIDInvalidError
		return
	}

	list, err := f.Service.GetFansList(uint(userId))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &ResponseWithData{
		Code: SuccessCode,
		Msg:  "ok",
		Data: list,
	})
}
