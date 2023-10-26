package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"utopia-back/database/implement"
	"utopia-back/service"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		Service: &service.UserService{
			Dal: &implement.UserImpl{},
		},
	}
}

type data struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserController) Login(c *gin.Context) {

	var r userRequest
	// 接收参数并绑定
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, &Response{
			Code: ErrorCode,
			Msg:  err.Error(),
		})
	}
	// 登录
	token, id, err := u.Service.Login(r.Username, r.Password)
	if err != nil {
		c.JSON(200, &Response{
			Code: ErrorCode,
			Msg:  err.Error(),
		})
		return
	}

	// 成功登录
	c.JSON(200, &Response{
		Code: SuccessCode,
		Msg:  "ok",
		Data: data{
			Token:  token,
			UserId: int(id),
		},
	})

}

func (u *UserController) Register(c *gin.Context) {

	var r userRequest
	// 接收参数并绑定
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, &Response{
			Code: ErrorCode,
			Msg:  err.Error(),
		})
	}
	// 注册
	token, id, err := u.Service.Register(r.Username, r.Password)
	if err != nil {
		c.JSON(200, &Response{
			Code: ErrorCode,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &Response{
		Code: SuccessCode,
		Msg:  "ok",
		Data: data{
			Token:  token,
			UserId: int(id),
		}})
}
