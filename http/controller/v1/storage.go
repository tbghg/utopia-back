package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"net/http"
	"strconv"
	"utopia-back/config"
	"utopia-back/http/controller/base"
	"utopia-back/pkg/logger"
	utils "utopia-back/pkg/util"
	"utopia-back/service/abstract"
)

type StorageController struct {
	StorageService abstract.StorageService
}

type uploadTokenData struct {
	Token string `json:"token"`
}

type uploadCallbackReq struct {
	Key         string `json:"key" validate:"required"`
	FileType    string `json:"file_type" validate:"required"`
	Uid         string `json:"uid"` // todo 可更改为JWT-Token，增强安全性
	VideoTypeId string `json:"video_type_id"`
	CoverUrl    string `json:"cover_url"`
	Describe    string `json:"describe"`
}

type callbackData struct {
	ImageUrl string `json:"image_url"`
}

const (
	callbackCover  = "COVER"  // 封面
	callbackAvatar = "AVATAR" // 头像
)

func (v *StorageController) UploadToken(c *gin.Context) {
	upToken := utils.GetCallbackToken()
	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: uploadTokenData{
			Token: upToken,
		},
	})
}

func (v *StorageController) UploadCallback(c *gin.Context) {
	var (
		r                uploadCallbackReq
		err              error
		uid, videoTypeId uint
	)

	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
		}
		logger.Logger.Error(fmt.Sprintf("UploadCallback err:%+v", err))
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
	uid, videoTypeId, err = callbackReqValidate(r)
	if err != nil {
		return
	}

	url := config.V.GetString("qiniu.kodoApi") + "/" + r.Key
	if r.FileType == callbackAvatar {
		err = v.StorageService.UpdateAvatar(uid, url)
		if err != nil {
			return
		}
	}
	if r.FileType == callbackAvatar || r.FileType == callbackCover {
		c.JSON(http.StatusOK, &base.ResponseWithData{
			Code: base.SuccessCode,
			Msg:  "ok",
			Data: callbackData{
				ImageUrl: url,
			},
		})
		return
	}
	err = v.StorageService.UploadVideoCallback(uid, url, r.CoverUrl, r.Describe, videoTypeId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &base.ResponseWithoutData{
		Code: base.SuccessCode,
		Msg:  "ok",
	})

}

func callbackReqValidate(r uploadCallbackReq) (uid uint, videoTypeId uint, err error) {
	if err = utils.Validate.Struct(r); err != nil {
		return
	}
	// 封面不做校验
	if r.FileType == callbackCover {
		return
	}
	aid, err1 := strconv.ParseUint(r.Uid, 10, 64)
	if err1 != nil {
		err = errors.New("uid参数传递错误")
		return
	}
	// 头像有uid即可，不用进一步校验
	if r.FileType == callbackAvatar {
		return uint(aid), 0, nil
	}

	if r.CoverUrl == "null" {
		err = errors.New("cover_url参数错误，请携带封面")
		return
	}

	tid, err1 := strconv.ParseUint(r.VideoTypeId, 10, 64)
	if err1 != nil {
		err = errors.New("video_type_id参数错误")
		return
	}
	uid, videoTypeId = uint(aid), uint(tid)
	return
}
