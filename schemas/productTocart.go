package schemas

import (
	"gorm.io/gorm"
)

type ProductsCarts struct {
	gorm.Model
	ProductID int `gorm:"primaryKey"`
	CartID    int `gorm:"primaryKey"`
	Quantity  int
}
