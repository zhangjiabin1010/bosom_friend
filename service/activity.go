package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActivityShow(c * gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"message":"获取成功",
	})
}
