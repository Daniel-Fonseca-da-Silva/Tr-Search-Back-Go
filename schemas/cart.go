package schemas

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Status   bool `gorm:"type:boolean;not null"`
	Products []ProductsCarts
	UserID   uint
}
