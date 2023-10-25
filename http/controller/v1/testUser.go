package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"utopia-back/database/implement"
	"utopia-back/service"
)

type TestUserController struct {
	Service *service.TestUserService
}

func NewTestUserCtrl() *TestUserController {
	return &TestUserController{
		Service: &service.TestUserService{
			Dal: &implement.TestUserImpl{},
		},
	}
}

func (t *TestUserController) Add(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return
	}
	id, err := t.Service.Add(name, ageInt)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (t *TestUserController) Select(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	user, err := t.Service.Select(uint(idInt))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"name": user.Name, "age": user.Age})
}
