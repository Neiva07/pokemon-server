package controllers

import (
	"net/http"
	"pokemon-server/models"
	u "pokemon-server/utils"
)

//GetAllAccounts return all accounts from the server
var GetAllAccounts = func(w http.ResponseWriter, r *http.Request) {

	accounts := &[]models.Account{}
	models.DB.GetDB().Table("accounts").Find(accounts)

	resp := u.Message(true, "All accounts register")

	resp["accounts"] = accounts

	u.Response(w, resp)
}
