package entity

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	CreateMigrations(database)
	if err != nil {
		return
	}
	DB = database
}

func getDsn() string {
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	return user + ":" + password + "@" + host + "/" + database + "?charset=utf8mb4"
}

func CreateMigrations(database *gorm.DB) {
	err := database.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
