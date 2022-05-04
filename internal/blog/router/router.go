package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ComradeProgrammer/blog/internal/blog/controller"
)

func GetGinEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	return r
}
