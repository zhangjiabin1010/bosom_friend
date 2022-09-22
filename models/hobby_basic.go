package models

import "time"

type Hobby struct {
	HobbyId     int       `gorm:"column:hobbyid"`
	Name        string    `gorm:"column:name"`
	Type        string    `gorm:"column:type"`
	Description string    `gorm:"column:description"`
	Createtime  time.Time `gorm:"column:createtime"`
}

func (table *Hobby) TableName() string {
	return "hobby_basic"
}