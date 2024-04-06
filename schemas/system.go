package schemas

import "gorm.io/gorm"

type System struct {
	gorm.Model
	About   string `gorm:"type:text;not null"`
	Title   string `gorm:"type:varchar(100);not null"`
	Version int    `gorm:"not null"`
	AdminID uint
}
