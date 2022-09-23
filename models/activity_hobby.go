package models

import "time"

type ActivityHobby struct {
	ActivityId int       `gorm:"column:activityid;primaryKey"`
	HobbyId    string    `gorm:"column:hobbyid;primaryKey"`
	Createtime time.Time `gorm:"column:createtime"`
}

func (table *ActivityHobby) TableName() string {
	return "activity_hobby"
}
