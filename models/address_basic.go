package models

type Address struct {
	Id       uint   `gorm:"column:id" db:"id" json:"id" form:"id"`
	Code     int64  `gorm:"column:code" db:"code" json:"code" form:"code"`                 //行政区划代码
	Name     string `gorm:"column:name" db:"name" json:"name" form:"name"`                 //名称
	Province string `gorm:"column:province" db:"province" json:"province" form:"province"` //省/直辖市
	City     string `gorm:"column:city" db:"city" json:"city" form:"city"`                 //市
	Area     string `gorm:"column:area" db:"area" json:"area" form:"area"`                 //区
	Town     string `gorm:"column:town" db:"town" json:"town" form:"town"`                 //城镇地区
}

func (address Address) TableName() string {
	return "address_basic"
}

func AddressQuery() {
	db := GetDB()
	db.Where()

}
