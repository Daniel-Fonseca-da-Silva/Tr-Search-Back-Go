package schemas

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Description string  `gorm:"type:text;"`
	Status      bool    `gorm:"type:boolean;not null"`
	Code        string  `gorm:"type:varchar(7);"`
	Price       float32 `gorm:"type:decimal(10,2);"`
	Address     []Address
	Photo       []Photo    `gorm:"many2many:products_to_photos;"`
	Category    []Category `gorm:"many2many:products_to_categories;"`
	Carts       []ProductsCarts
}
