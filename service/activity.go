package service

import (
	"bosom_friend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ActivityQuery
// @Tags 公共方法
// @Summary 活动查询
// @Param page query int false "当前页数:默认第一页"
// @Param size query int false "每页数据量"
// @Param hobbyid query int false "爱好id"
// @Param locationid query int false "地区id"
// @Success 200 {string} json "{"code":200, "message":"获取成功", "data":"",}"
// @Router /activityquery [get]
func ActivityQuery(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	hobbyid := c.Query("hobbyid")
	locationid := c.Query("locationid")
	offset := (page - 1) * size
	fmt.Println(offset)
	activitysData := models.GetActivityList(hobbyid, locationid)
	fmt.Println(activitysData)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    "",
	})
	//c.JSON(http.StatusOK,gin.H{
	//	"code":200,
	//	"message":"获取成功",
	//	"data": map[string] interface{} {
	//		"data":data,
	//		"count":count,
	//	},
	//})
}

func ActivityAdd(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "增加成功",
	})
}

func ActivityUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

func ActivityDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
