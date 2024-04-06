package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"type:varchar(50);not null"`
	Password      string `gorm:"type:varchar(20);not null"`
	DisplayName   string `gorm:"type:varchar(20);not null"`
	Address       Address
	Admin         []Admin
	Photo         []Photo
	Favorite      Favorite
	Cart          []Cart
	Configuration Configuration
}
