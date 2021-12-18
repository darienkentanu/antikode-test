package util

import (
	"antikode-test/config"
	"antikode-test/models"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDatabaseConnection(config *config.AppConfig) *gorm.DB {
	// var uri string

	var uri = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username, config.Database.Password,
		config.Database.Address, config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Info("failed connecting to the database: ", err)
		panic(err)
	}

	DatabaseMigration(db)
	return db
}

func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Brand{})
	db.AutoMigrate(&models.Outlet{})
	db.AutoMigrate(&models.Product{})
}
