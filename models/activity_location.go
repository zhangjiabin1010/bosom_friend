package models

import "time"

type ActivityLocation struct {
	ActivityId int         `gorm:"column:activityid;primaryKey"`
	Code       string      `gorm:"column:code;primaryKey"`
	Createtime time.Time   `gorm:"column:createtime"`
	Activityid []*Activity `gorm:"many2many:activity_location;"`
}

func (table *ActivityLocation) TableName() string {
	return "activity_location"
}
