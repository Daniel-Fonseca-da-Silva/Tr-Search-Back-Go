package schemas

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:TEXT;not null"`
	Active      bool   `gorm:"type:boolean;not null"`
	VehicleID   uint
	HouseID     uint
}
