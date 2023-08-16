package conn

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ComradeProgrammer/blog/internal/myblog/dal/model"
)

var WatchList = []interface{}{
	&model.Blog{},
	&model.Category{},
	&model.User{},
	&model.Comment{},
}

func ConnectSqliteDatabase(databaseName string) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{
		Logger: logger.Discard, //shutdown output
	})
	if err != nil {
		log.Printf("failed to open sqlite db %s: %v\n", databaseName, err)
		return nil, err
	}

	//enable foreign key check
	database.Exec("PRAGMA foreign_keys = ON;")
	if database.Error != nil {
		log.Printf("failed to enable foreign key: %v\n", err)
		return nil, err
	}
	// start automigration
	for _, i := range WatchList {
		if err := database.AutoMigrate(i); err != nil {
			log.Printf("failed to migrate foreign key %v: %v\n", i, err)
			return nil, err

		}
	}
	return database, nil
}

