package router

import (
	"bosom_friend/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	r := gin.Default()
	r.GET("/activityshow",service.ActivityShow)
	return r
}