package router

import (
	_ "bosom_friend/docs"
	"bosom_friend/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/activityquery", service.ActivityQuery)
	r.POST("/activityadd", service.ActivityAdd)
	r.POST("/activityupdate", service.ActivityUpdate)
	r.DELETE("/activitydelete", service.ActivityDelete)

	return r
}
