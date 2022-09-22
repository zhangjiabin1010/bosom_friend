package models

import "time"

type UserHobby struct {
	Id         int       `gorm:"column:id"`
	HobbyId    string    `gorm:"column:hobbyid"`
	Createtime time.Time `gorm:"column:createtime"`
}

func (table *UserHobby) TableName() string {
	return "user_hobby"
}
