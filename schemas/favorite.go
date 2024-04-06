package schemas

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	DisplayName string  `gorm:"type:varchar(30);not null"`
	Opinion     string  `gorm:"type:varchar(30);not null"`
	Pontuation  float64 `gorm:"not null"`
	SellerID    uint
	UserID      uint
}
