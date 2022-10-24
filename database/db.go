package database

import (
	"apigolang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var InstanceData *gorm.DB

var databaseError error

func ConnectDB(connectionString string) () {
	InstanceData, databaseError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if databaseError != nil {
		log.Fatal(databaseError)
		panic("Tidak bisa koneksi ke database")
	}
	log.Println("Koneksi ke database berhasil")
}

func Migrate() {
	InstanceData.AutoMigrate(&models.User{}, &models.Task{})
	log.Println("Database Migration Successfully")
}