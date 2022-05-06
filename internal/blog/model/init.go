package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

var WatchList = []any{
	&Blog{},
	&Category{},
	&User{},
}

func ConnectDatabase(databaseName string) {
	var err error
	database, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{
		Logger: logger.Discard, //shutdown output
	})
	if err != nil {
		panic("Error: failed to connect database")
	}
	//enable foreign key check
	database.Exec("PRAGMA foreign_keys = ON;")

	for _, i := range WatchList {
		database.AutoMigrate(i)
	}

}

func InitDatabase() {
	//insert the default admin account

	adminUser, err := GetUserByUserName("admin")

	if err == gorm.ErrRecordNotFound {
		adminUser = &User{
			UserName: "admin",
			IsAdmin:  true,
		}
		adminUser.SetPassword("123456")
		err = CreateUser(adminUser)
		if err != nil {
			log.Println("Unable to create admin account")
		}
	}
}

func ClearDatabase() {
	//if we do not se 1=1 conditon, gorm will reject delete operation with no conditions
	for _, i := range WatchList {
		database.Where("1 = 1").Delete(i)
	}
}
