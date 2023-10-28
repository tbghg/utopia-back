package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"utopia-back/service/implement/v1"
)

type TestUserController struct {
	TestUserService *v1.TestUserService
}

func NewTestUserCtrl() *TestUserController {
	return &TestUserController{
		TestUserService: v1.NewTestUserService(),
	}
}

func (t *TestUserController) Add(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return
	}
	id, err := t.TestUserService.Add(name, ageInt)
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
	user, err := t.TestUserService.Select(uint(idInt))
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"name": user.Name, "age": user.Age})
}
