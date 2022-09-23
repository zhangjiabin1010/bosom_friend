package models

import (
	"fmt"
	"time"
)

type Activity struct {
	Id           int         `gorm:"column:id"`
	ActivityId   string      `gorm:"column:activityid"`
	Title        string      `gorm:"column:title"`
	Condition    string      `gorm:"column:condition"`
	Number       string      `gorm:"column:number"`
	Gender       string      `gorm:"column:gender"`
	Age          string      `gorm:"column:age"`
	ActivityTime time.Time   `gorm:"column:activitytime"`
	Status       string      `gorm:"column:status"`
	Createtime   time.Time   `gorm:"column:createtime"`
	Updatetime   time.Time   `gorm:"column:updatetime"`
	Locations    []*Location `gorm:"many2many:activity_location"`
	Hobbys       []*Hobby    `gorm:"many2many:activity_hobby"`
}

func (table *Activity) TableName() string {
	return "activity_basic"
}

func GetActivityList(hobbyid string, locationid string) (int64, []Activity) {
	tx := GetDB()
	var count int64
	var activitys []Activity
	tx.Model(&Activity{}).Where("number = 2").Count(&count).Scan(&activitys)
	//data := make(map[string]interface{})
	//data["count"] = count
	//data["data"] = activitys
	return count, activitys

}

func GetActivityTest(hobbyid string, locationid string) (int64, []Activity) {
	tx := GetDB()
	var count int64
	var activitys []Activity
	//tx.Model(&Activity{}).Preload("models.Hobby").Find(&activitys)
	err := tx.Model(&Activity{}).Association("Hobbys").Find(&activitys)
	if err != nil {
		fmt.Println("OK")
	}

	//data := make(map[string]interface{})
	//data["count"] = count
	//data["data"] = activitys
	return count, activitys

}
