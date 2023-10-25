package main

import (
	"github.com/gin-gonic/gin"
	"utopia-back/initialize"
)

func main() {
	r := gin.New()
	initialize.InitRoute(r)
	r.Run()
}
