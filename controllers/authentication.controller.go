package controllers

import (
	"encoding/json"
	"net/http"
	"pokemon-server/models"
	u "pokemon-server/utils"
)

// Token to represent a token of a user

// Authenticate check and auth the user who is trying to signin
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		u.Response(w, u.Message(false, "Invalid field. Please, try again"))
		return
	}

	response := models.Login(account.Email, account.Password)
	u.Response(w, response)

}

// SignUp create a user in the database
var SignUp = func(w http.ResponseWriter, r *http.Request) {

	newAccount := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(newAccount)

	if err != nil {
		u.Response(w, u.Message(false, "Something went wrong. Please, try again."))
		return
	}

	response := newAccount.Create()

	u.Response(w, response)
}
