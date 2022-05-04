package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDatabase(databaseName string) {
	var err error
	database, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		panic("Error: failed to connect database")
	}
	//enable foreign key check
	database.Exec("PRAGMA foreign_keys = ON;")

	database.AutoMigrate(&Blog{})
	database.AutoMigrate(&Category{})
}

func ClearDatabase() {
	database.Where("1 = 1").Delete(&Blog{})
	database.Where("1 = 1").Delete(&Category{})
}
