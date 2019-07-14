package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//Database the data struct to hold the db itself and the methods attached to it
type Database struct {
	db *gorm.DB
}

//Initialize method to initialize the database
func (DB *Database) Initialize() {

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
	DB.db = conn
	DB.db.Debug().AutoMigrate(&Account{})

}

//GetDB --exporting the database to the rest of the application
func (DB *Database) GetDB() *gorm.DB {
	return DB.db
}

//DB declare the variable that could be exported
var DB Database

func init() {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Print(err)
	// }
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/pokemon-server/.env"))

	if err != nil {
		log.Print(err)
	}
	DB.Initialize()
}
