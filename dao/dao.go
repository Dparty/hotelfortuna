package dao

import (
	"fmt"
	"god-of-wealth/common/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		user := config.GetString("database.user")
		password := config.GetString("database.password")
		host := config.GetString("database.host")
		port := config.GetString("database.port")
		database := config.GetString("database.database")
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, database,
		)
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}
	return db
}
