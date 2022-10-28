package models

import (
	"gorm.io/gorm"
	"time"
)

type Hobby struct {
	gorm.Model `json:"-"`
	HobbyId     int         `gorm:"column:hobbyid"`
	Name        string      `gorm:"column:name"`
	Type        string      `gorm:"column:type"`
	Description string      `gorm:"column:description"`
	Createtime  time.Time   `gorm:"column:createtime"`
	Activitys   []*Activity `gorm:"many2many:activity_hobby"`
}

func (table *Hobby) TableName() string {
	return "hobby_basic"
}
