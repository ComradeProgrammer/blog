package router

import (
	"github.com/gin-gonic/gin"

	"github.com/ComradeProgrammer/blog/internal/blog/controller"
)

func StartBackendServer() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run()
}
