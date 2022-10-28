package models

import (
	"fmt"
	"time"
)


type Activity struct {
	Id           int         `gorm:"column:id"`
	ActivityId   string      `gorm:"column:activityid"`
	Title        string      `gorm:"column:title"`
	Conditions    string      `gorm:"column:conditions"`
	Number       string      `gorm:"column:number"`
	Gender       string      `gorm:"column:gender"`
	Age          string      `gorm:"column:age"`
	ActivityTime time.Time   `gorm:"column:activitytime"`
	Status       string      `gorm:"column:status"`
	Createtime   time.Time   `gorm:"column:createtime"`
	Updatetime   time.Time   `gorm:"column:updatetime"`
	//Locations    []*Location `gorm:"many2many:activity_location"`
	Hobbys       []*Hobby    `gorm:"many2many:activity_hobby"`


}

func (table *Activity) TableName() string {
	return "activity_basic"
}

func GetActivityList(hobbyid string, locationid string) (int64, []Activity) {
	tx := GetDB()
	var count int64
	var activitys []Activity
	tx.Model(&Activity{}).
		Joins("inner join activity_hobby ah on ah.activityid = activity_basic.activityid").
		Where("activity_basic.hobbyid = ?","1").
		Select("activityid,title,condition").
		Count(&count).Find(&activitys)
	//data := make(map[string]interface{})
	//data["count"] = count
	//data["data"] = activitys
	return count, activitys

}

func GetActivityTest(hobbyid string, locationid string) (int64, []Activity) {
	tx := GetDB()

	var count int64
	var activitys []Activity
	var hobbys []Hobby
	//tx.Model(&Activity{}).Preload("models.Hobby").Find(&activitys)
	//err := tx.Model(&Activity{}).
	//	Joins("inner join activity_hobby ah on ah.hobbyid = hobby.hobbyid").
	//	Select(activity.activityid,)
	data := tx.Model(&Activity{}).
		Joins("inner join activity_hobby ah on ah.activityid = activity_basic.activityid and ah.hobbyid = ?","1").
		Select("activity_basic.activityid,activity_basic.title,activity_basic.conditions").
		Count(&count).Find(&activitys)


	//data := make(map[string]interface{})
	//data["count"] = count
	//data["data"] = activitys
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxx")
	fmt.Println(data)
	fmt.Println(activitys)
	fmt.Println(hobbys)
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxx")

	return count, activitys

}
