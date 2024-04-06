package schemas

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Filename string `gorm:"type:varchar(7);"`
	Url      string `gorm:"type:varchar(255)"`
	UserID   uint
	Seller   Seller
}
