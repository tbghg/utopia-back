package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"net/http"
	"strconv"
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
	IsImage  string `json:"is_image"`
	AuthorId string `json:"author_id"` // todo 可更改为JWT-Token，增强安全性
	CoverUrl string `json:"cover_url"`
	Describe string `json:"describe"`
}

type uploadCallbackData struct {
	ImageUrl string `json:"image_url"`
}

// 上传是否为图片
const callbackIsImage = "YES"

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
	// 校验是否为七牛云调用
	isQiNiu, err := qbox.VerifyCallback(utils.GetMac(), c.Request)
	if !isQiNiu || err != nil {
		if err == nil {
			err = errors.New("非七牛云服务发送")
		}
		return
	}
	// 接收参数并绑定
	if err = c.BindJSON(&r); err != nil {
		return
	}
	// 参数校验
	if err = utils.Validate.Struct(r); err != nil {
		return
	}

	url := config.V.GetString("qiniu.kodoApi") + "/" + r.Key
	if r.IsImage == callbackIsImage {
		c.JSON(http.StatusOK, &ResponseWithData{
			Code: SuccessCode,
			Msg:  "ok",
			Data: uploadCallbackData{
				ImageUrl: url,
			},
		})
		return
	}
	authorId, _ := strconv.ParseUint(r.AuthorId, 10, 64)
	err = v.Service.UploadVideoCallback(uint(authorId), url, r.CoverUrl, r.Describe)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &ResponseWithoutData{
		Code: SuccessCode,
		Msg:  "ok",
	})

}
