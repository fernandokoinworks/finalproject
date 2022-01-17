package database

import (
	"finalproject/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Setup() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading  .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	PORT := os.Getenv("PORT")

	db, err := gorm.Open("postgres", "host="+host+
		" port = "+PORT+
		" user="+username+
		" dbname="+database+
		" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println("Success connect to database")
	db.AutoMigrate([]models.Todo{})
	DB = db

}
func GetDB() *gorm.DB {
	return DB
}
