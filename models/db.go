package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Print(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	log.Println(dbURI)
	conn, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Print(err)
	}
	db = conn
	db.Debug().AutoMigrate(&Account{})

}

//GetDB --exporting the database to the rest of the application
func GetDB() *gorm.DB {
	return db
}
