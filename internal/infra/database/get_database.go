package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDatabase(entity interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(entity)

	if err != nil {
		return nil, err
	}

	return db, nil
}
