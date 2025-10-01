package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func getENV(key, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultValue
	}
	return env
}

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() (err error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=ut?f8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		getENV("DB_USER", "root"),
		getENV("DB_PASSWORD", "password"),
		getENV("DB_HOST", "localhost"),
		getENV("DB_PORT", "3306"),
		getENV("DB_NAME", "CineVerse"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})
	return err
}

func Get() *gorm.DB {
	return db
}
