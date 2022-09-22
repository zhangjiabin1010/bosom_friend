package models

import "time"

type User struct {
	Id         int       `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Password   string    `gorm:"column:password"`
	Gender     string    `gorm:"column:gender"`
	Birth      string    `gorm:"column:birth"`
	Age        string    `gorm:"column:age"`
	Credit     string    `gorm:"column:credit"`
	Coin       string    `gorm:"column:coin"`
	Phone      string    `gorm:"column:phone"`
	Source     string    `gorm:"column:source"`
	Status     string    `gorm:"column:status"`
	Wx_openid  string    `gorm:"column:wx_openid"`
	Createtime time.Time `gorm:"column:createtime"`
}

func (table *User) TableName() string {
	return "user_basic"
}
