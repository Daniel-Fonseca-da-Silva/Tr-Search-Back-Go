package schemas

import (
	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	Energy        float64
	Water         float64
	Kitchen       int  `gorm:"not null"`
	Furniture     bool `gorm:"type:boolean;not null"`
	Level         int
	Garden        int
	Animals       bool `gorm:"type:boolean;not null"`
	Bathroom      int  `gorm:"not null"`
	Room          int  `gorm:"not null"`
	Hall          int  `gorm:"not null"`
	Balcony       int
	Garage        int `gorm:"not null"`
	Size          int
	Guarantor     bool    `gorm:"type:boolean;not null"`
	EntranceValue float64 `gorm:"not null"`
	DepositValue  float64 `gorm:"not null"`
	Pool          int
	Category      []Category
}
