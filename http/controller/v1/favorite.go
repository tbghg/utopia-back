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
	UserId int `json:"user_id" validate:"required,gt=0"` // 用户id > 0 必需
	Video  int `json:"video" validate:"required,gt=0"`   // 视频id > 0 必需
}

// AddFavorite 添加收藏
func (f *FavoriteController) AddFavorite(c *gin.Context) {
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

	// 接收参数并绑定
	if err = c.ShouldBindJSON(&r); err != nil {
		return
	}
	// 参数校验
	if err = utils.Validate.Struct(r); err != nil {
		return
	}
	// 添加收藏
	err = f.FavoriteService.AddFavorite(uint(r.UserId), uint(r.Video))
	if err != nil {
		return
	}

	// 成功添加收藏
	c.JSON(200, SuccessResponse)

}

// DeleteFavorite 删除收藏
func (f *FavoriteController) DeleteFavorite(c *gin.Context) {
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

	// 接收参数并绑定
	if err = c.ShouldBindJSON(&r); err != nil {
		return
	}
	// 参数校验
	if err = utils.Validate.Struct(r); err != nil {
		return
	}
	// 删除收藏
	err = f.FavoriteService.DeleteFavorite(uint(r.UserId), uint(r.Video))
	if err != nil {
		return
	}
	// 成功删除收藏
	c.JSON(200, SuccessResponse)
}
