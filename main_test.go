package main

import (
	"log"
	"os"
	"pokemon-server/models"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {

	db := models.DB.GetDB()
	defer clearTable(db)

	app := App{}
	app.UseRouter()

	os.Exit(m.Run())

}
func clearTable(db *gorm.DB) {
	db.Exec("DELETE FROM accounts;")
	db.Exec("ALTER SEQUENCE accounts_id_seq RESTART WITH 1;")
	accounts := &[]models.Account{}
	models.DB.GetDB().Table("accounts").Find(accounts)

	log.Print(*accounts)
	log.Print("Done!")
}
