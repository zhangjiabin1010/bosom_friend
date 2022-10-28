package service

import (
	"bosom_friend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ActivityQuery
// @Tags 地址方法
// @Summary 地址信息查询
// @Success 200 {string} json "{"code":200, "message":"获取成功", "data":"",}"
// @Router /addressquery [get]

func AddressQuery(c *gin.Context) {
	count, activitysData := models.GetActivityList(hobbyid, locationid)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": map[string]interface{}{
			"data":  activitysData,
			"count": count,
		},
	})

}