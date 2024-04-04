package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Password    string
	DisplayName string
}
