package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

var WatchList = []any{
	&Blog{},
	&Category{},
	&User{},
}

func InitDatabase(databaseName string) {
	var err error
	database, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		panic("Error: failed to connect database")
	}
	//enable foreign key check
	database.Exec("PRAGMA foreign_keys = ON;")

	for _, i := range WatchList {
		database.AutoMigrate(i)
	}
}

func ClearDatabase() {
	//if we do not se 1=1 conditon, gorm will reject delete operation with no conditions
	for _, i := range WatchList {
		database.Where("1 = 1").Delete(i)
	}
}
