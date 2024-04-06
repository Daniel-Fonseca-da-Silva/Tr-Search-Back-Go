package schemas

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Documentation string `gorm:"type:varchar(30);not null"`
	UserID        uint
	System        System
}
