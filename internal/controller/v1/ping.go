package controller_v1

import (
	"github.com/gin-gonic/gin"
)

//PingController ...
type PingController struct {
}

//Ping ...
func (c *PingController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

//NewPingController ...
func NewPingController() *PingController {
	return &PingController{}
}
