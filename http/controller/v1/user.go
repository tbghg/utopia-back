package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	utils "utopia-back/pkg/util"
	"utopia-back/service/abstract"
)

type UserController struct {
	Service abstract.UserService
}

func NewUserController(s abstract.UserService) *UserController {
	return &UserController{Service: s}
}

type authData struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

type userRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserController) Login(c *gin.Context) {

	var (
		r   userRequest
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

	// 接收参数并绑定
	if err = c.ShouldBindJSON(&r); err != nil {
		return
	}
	// 参数校验
	if err = utils.Validate.Struct(r); err != nil {
		return
	}
	// 登录
	token, id, err := u.Service.Login(r.Username, r.Password)
	if err != nil {
		return
	}

	// 成功登录
	c.JSON(200, &ResponseWithData{
		Code: SuccessCode,
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
			c.JSON(http.StatusOK, &ResponseWithData{
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
	// 注册
	token, id, err := u.Service.Register(r.Username, r.Password)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, &ResponseWithData{
		Code: SuccessCode,
		Msg:  "ok",
		Data: authData{
			Token:  token,
			UserId: int(id),
		}})
}
