package db

import (
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/error-messages"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := configs.GetEnv().DBConstraints
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(error_messages.Error.DbConnectionError)
	}
	DB = db
	fmt.Println("DB connection done")
}
