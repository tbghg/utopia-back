package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"utopia-back/config"
	utils "utopia-back/pkg/util"
	"utopia-back/service/abstract"
)

type VideoController struct {
	Service abstract.VideoService
}

func NewVideoController(s abstract.VideoService) *VideoController {
	return &VideoController{Service: s}
}

type uploadVideoTokenData struct {
	Token string `json:"token"`
}

type uploadCallbackReq struct {
	Key      string `json:"key" validate:"required"`
	IsImage  bool   `json:"is_image"`
	AuthorId uint   `json:"author_id" `
	CoverUrl string `json:"cover_url" `
	Describe string `json:"describe"`
}

type uploadCallbackData struct {
	ImageUrl string `json:"image_url"`
}

func (v *VideoController) UploadVideoToken(c *gin.Context) {
	upToken := utils.GetCallbackToken()
	c.JSON(http.StatusOK, &ResponseWithData{
		Code: SuccessCode,
		Msg:  "ok",
		Data: uploadVideoTokenData{
			Token: upToken,
		},
	})
}

func (v *VideoController) UploadVideoCallback(c *gin.Context) {
	var (
		r   uploadCallbackReq
		err error
	)

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &ResponseWithData{
				Code: ErrorCode,
				Msg:  err.Error(),
			})
		}
	}()
	// todo 校验是否为七牛云发送

	// 接收参数并绑定
	if err = c.ShouldBindJSON(&r); err != nil {
		return
	}
	// 参数校验
	if err = utils.Validate.Struct(r); err != nil {
		return
	}

	url := config.V.GetString("qiniu.kodoApi") + r.Key
	if r.IsImage {
		c.JSON(http.StatusOK, &ResponseWithData{
			Code: SuccessCode,
			Msg:  "ok",
			Data: uploadCallbackData{
				ImageUrl: url,
			},
		})
		return
	}

	err = v.Service.UploadVideoCallback(r.AuthorId, url, r.CoverUrl, r.Describe)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &ResponseWithoutData{
		Code: SuccessCode,
		Msg:  "ok",
	})

}
