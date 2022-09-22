package models

import "time"

type ActivityHobby struct {
	ActivityId int       `gorm:"column:activityid"`
	HobbyId    string    `gorm:"column:hobbyid"`
	Createtime time.Time `gorm:"column:createtime"`
}

func (table *ActivityHobby) TableName() string {
	return "activity_hobby"
}
