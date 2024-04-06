package schemas

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model
	Font     string `gorm:"type:varchar(20);not null"`
	Theme    string `gorm:"type:varchar(20);not null"`
	TypeUser string `gorm:"type:varchar(20);not null"`
	SellerID uint
	UserID   uint
}
