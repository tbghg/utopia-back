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

type UserController struct {
	UserService abstract.UserService
}

type authData struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

type userRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type updateNicknameReq struct {
	Nickname string `json:"nickname" validate:"required"`
}

func (u *UserController) Login(c *gin.Context) {

	var (
		r   userRequest
		err error
	)
	// 请求处理失败，返回错误信息
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("Login err:%+v", err))
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
	// 登录
	token, id, err := u.UserService.Login(r.Username, r.Password)
	if err != nil {
		return
	}

	// 成功登录
	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: authData{
			Token:  token,
			UserId: int(id),
		},
	})

}

func (u *UserController) Register(c *gin.Context) {

	var (
		r   userRequest
		err error
	)

	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("Register err:%+v", err))
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
	// 注册
	token, id, err := u.UserService.Register(r.Username, r.Password)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: authData{
			Token:  token,
			UserId: int(id),
		}})
}

func (u *UserController) UpdateNickname(c *gin.Context) {
	var (
		err error
		r   updateNicknameReq
	)
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("UpdateNickname err:%+v", err))
		}
	}()

	// 获取用户id
	value, ok := c.Get("user_id")
	uid, ok := value.(int)
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
	err = u.UserService.UpdateNickname(uint(uid), r.Nickname)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, base.SuccessResponse)
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusOK, &base.ResponseWithoutData{
				Code: base.ErrorCode,
				Msg:  err.Error(),
			})
			logger.Logger.Error(fmt.Sprintf("UpdateNickname err:%+v", err))
		}
	}()

	// 获取用户id
	value, ok := c.Get("user_id")
	uid, ok := value.(int)
	if !ok {
		err = base.UserIDInvalidError
		return
	}

	info, err := u.UserService.GetUserInfo(uint(uid))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &base.ResponseWithData{
		Code: base.SuccessCode,
		Msg:  "ok",
		Data: info,
	})

	return
}
