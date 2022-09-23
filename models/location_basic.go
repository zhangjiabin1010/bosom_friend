package models

type Location struct {
	Id        int         `gorm:"column:id"`
	Code      string      `gorm:"column:code"`
	Name      string      `gorm:"column:name"`
	Province  string      `gorm:"column:province"`
	City      string      `gorm:"column:city"`
	Area      string      `gorm:"column:area"`
	Town      string      `gorm:"column:town"`
	Activitys []*Activity `gorm:"many2many:activity_location"`
}

func (table *Location) TableName() string {
	return "location_basic"
}
