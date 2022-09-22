package models

import (
	"time"
)

type Activity struct {
	Id           int       `gorm:"column:id"`
	ActivityId   string    `gorm:"column:activityid"`
	Title        string    `gorm:"column:title"`
	Condition    string    `gorm:"column:condition"`
	Number       string    `gorm:"column:number"`
	Gender       string    `gorm:"column:gender"`
	Age          string    `gorm:"column:age"`
	Activitytime time.Time `gorm:"column:activitytime"`
	Status       string    `gorm:"column:status"`
	Createtime   time.Time `gorm:"column:createtime"`
	Updatetime   time.Time `gorm:"column:updatetime"`
}

func (table *Activity) TableName() string {
	return "activity_basic"
}

func GetActivityList(hobbyid string, locationid string) interface{} {
	tx := GetDB()
	var count int64
	var activitys []Activity
	tx.Model(&Activity{}).Count(&count).Scan(&activitys)
	data := make(map[string]interface{})
	data["count"] = count
	data["data"] = activitys
	return data

}
