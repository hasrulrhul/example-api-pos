package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@(127.0.0.1:3308)/api_pos?charset=utf8mb4&parseTime=True&loc=Local"
	config := gorm.Config{}

	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		panic("Failed to connect to database!" + os.Getenv("DB_USER"))
	}
	DB = db

	// dsn := "root:root@tcp(127.0.0.1:3308)/api_pos?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("Failed to connect to database!" + os.Getenv("DB_USER"))
	// }
	// DB = db

}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
