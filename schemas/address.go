package schemas

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	City       string `gorm:"type:varchar(100);not null"`
	State      string `gorm:"type:char(2);not null"`
	Country    string `gorm:"type:char(2);not null"`
	PostalCode string `gorm:"type:varchar(30);not null"`
	Email      string `gorm:"type:varchar(60);not null"`
	Phone      string `gorm:"type:varchar(20);not null"`
	Address    string `gorm:"type:varchar(255);not null"`
	FirstName  string `gorm:"type:varchar(100);not null"`
	LastName   string `gorm:"type:varchar(100);not null"`
	Telephone  string `gorm:"type:varchar(20);not null"`
	UserID     uint
	ProductID  uint
	Seller     []Seller
}
