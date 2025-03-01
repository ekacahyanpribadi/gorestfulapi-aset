package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//db local
	//database, err := gorm.Open(mysql.Open("c5sWttLvLZ:hF1zJCDbvsVkodiMMzaq@tcp(127.0.0.1:3306)/easet"), &gorm.Config{})

	//db on premis
	database, err := gorm.Open(mysql.Open("c5sWttLvLZ:hF1zJCDbvsVkodiMMzaq@tcp(103.130.164.250:3306)/easet"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	//database.AutoMigrate(&Category{})
	DB = database

}
