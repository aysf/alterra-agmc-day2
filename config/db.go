package config

import (
	"day2/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var dbURL string = "root:custompwd@tcp(localhost:6033)/agmc?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("error opening db connection: ", err)
	}

	InitMigrate()
}

func InitMigrate() {

	DB.Migrator().DropTable(&models.User{}, &models.Book{})

	DB.AutoMigrate(&models.User{}, &models.Book{})

	var users = []models.User{
		{Name: "Adi", Email: "adi@example.com", Password: "pass123"},
		{Name: "Budi", Email: "budi@example.com", Password: "pass123"},
	}

	DB.Create(&users)

	var books = []models.Book{
		{Title: "Pulang", Author: "Tere Liye"},
		{Title: "Doraemon", Author: "Fujiko"},
		{Title: "Dragon Ball", Author: "Akira Toriyama"},
		{Title: "Kambing Jantan", Author: "Raditya Dika"},
	}

	DB.Create(&books)

}
