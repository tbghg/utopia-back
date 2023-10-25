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

func (u *UserController) Login(c *gin.Context) {

	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var r LoginRequest
	// 接收参数并绑定
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}
	// 登录
	token, id, err := u.Service.Login(r.Username, r.Password)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token, "id": id})
}

func (u *UserController) Register(c *gin.Context) {
	type RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var r RegisterRequest
	// 接收参数并绑定
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}
	// 注册
	token, id, err := u.Service.Register(r.Username, r.Password)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token, "id": id})
}
