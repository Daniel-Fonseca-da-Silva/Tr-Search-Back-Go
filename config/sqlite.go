package config

import (
	"os"

	"github.com/Daniel-Fonseca-da-Silva/Tr-Search-Back-Go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")
		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schemas.User{}, &schemas.Admin{}, &schemas.Address{}, &schemas.Product{}, &schemas.Photo{}, &schemas.Category{}, schemas.Vehicle{}, schemas.House{}, schemas.Favorite{}, schemas.Seller{}, schemas.Cart{}, schemas.ProductsCarts{}, schemas.Configuration{}, schemas.System{})
	if err != nil {
		logger.Errorf("Sqlite automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
