package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	utils "utopia-back/pkg/util"
	"utopia-back/service/abstract"
)

type FavoriteController struct {
	FavoriteService abstract.FavoriteService
}

func NewFavoriteController(service abstract.FavoriteService) *FavoriteController {
	return &FavoriteController{FavoriteService: service}
}

type favoriteRequest struct {
	VideoId    int `json:"video_id" validate:"required,gt=0"` // 视频id > 0 必需
	ActionType int `json:"action_type" validate:"required"`   // 操作类型 1: 添加收藏 2: 取消收藏
}

// Favorite 添加/取消 收藏
func (f *FavoriteController) Favorite(c *gin.Context) {
	var (
		r   favoriteRequest
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
	userID, ok := value.(int)
	if !ok {
		err = UserIDInvalidError
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
	// 判断操作类型
	if r.ActionType != 1 && r.ActionType != 2 {
		err = ActionTypeInvalidError
		return
	}

	if r.ActionType == 1 {
		err = f.FavoriteService.AddFavorite(uint(userID), uint(r.VideoId))
	} else {
		err = f.FavoriteService.DeleteFavorite(uint(userID), uint(r.VideoId))
	}

	// 成功添加收藏
	c.JSON(200, SuccessResponse)

}