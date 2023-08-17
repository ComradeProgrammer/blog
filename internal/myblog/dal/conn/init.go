package conn

import (
	"fmt"
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

	// if there is no admin, add an admin
	user := model.User{}
	result := database.Where("user_name = ?", "admin").First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		adminUser := &model.User{
			UserName: "admin",
			IsAdmin:  true,
		}
		adminUser.SetPassword("123456")
		result = database.Create(adminUser)
		if result.Error != nil {
			return nil, fmt.Errorf("unable to create admin account:%v ", err)
		}
	}
	return database, nil
}
