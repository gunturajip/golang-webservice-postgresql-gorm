package database

import (
	"fmt"
	"golang-webservice/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("Connecting to database...")
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}
	fmt.Println("Successfully connected to database🔥")

	db.Debug().AutoMigrate(models.User{}, models.Book{})
	CreateUser()
}

func GetDB() *gorm.DB {
	return db
}

func CreateUser() {
	db := GetDB()

	User := models.User{
		Email: "gunturaji852@gmail.com",
	}

	err := db.Create(&User).Error
	if err != nil {
		fmt.Println("Error Creating User Data :", err)
		return
	}
	fmt.Println("Created User Data")
}
