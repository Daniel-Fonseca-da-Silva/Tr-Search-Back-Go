package schemas

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	Site          string `gorm:"type:varchar(7);"`
	Document      string `gorm:"type:varchar(7);not null"`
	About         string `gorm:"type:varchar(7);not null"`
	Status        bool   `gorm:"type:boolean;not null"`
	Username      string `gorm:"type:varchar(50);not null"`
	Password      string `gorm:"type:varchar(20);not null"`
	DisplayName   string `gorm:"type:varchar(20);not null"`
	PhotoID       uint
	Favorite      Favorite
	AddressID     uint
	Configuration Configuration
}
