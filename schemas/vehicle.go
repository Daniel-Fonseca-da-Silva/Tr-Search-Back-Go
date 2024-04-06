package schemas

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Brand    string         `gorm:"type:varchar(30);not null"`
	Color    string         `gorm:"type:char(2);not null"`
	Year     datatypes.Date `gorm:"not null"`
	Door     int            `gorm:"not null"`
	Hp       int            `gorm:"not null"`
	Km       int            `gorm:"not null"`
	Category []Category
}
