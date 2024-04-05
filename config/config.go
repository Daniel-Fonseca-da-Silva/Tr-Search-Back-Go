package config

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!1")
	return errors.New("fake error")
	// return nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
