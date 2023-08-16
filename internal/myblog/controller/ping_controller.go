package controller

import "github.com/gin-gonic/gin"

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}
func (p *PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
