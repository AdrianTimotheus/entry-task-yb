package db

import (
	"fmt"
	"log"
	"os"

	"awesomeProject/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() (err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		helper.GetENV("DB_HOST", "localhost"),
		helper.GetENV("DB_USER", "root"),
		helper.GetENV("DB_PASSWORD", "password"),
		helper.GetENV("DB_NAME", "CineVerse"),
		helper.GetENV("DB_PORT", "5432")) // Default PostgreSQL port is 5432

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{ // Use postgres.Open()
		Logger: getLogger(),
	})
	return err
}

func Get() *gorm.DB {
	return db
}
