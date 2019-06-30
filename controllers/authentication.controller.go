package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/dgrijalva/jwt-go"
)

// Token to represent a token of a user
type Token struct {
	UserId uint
	jwt.StandardClaims
}

// Account struct to represent the account of a user
type Account struct {
	gorm.Model

	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Authenticate check and auth the user who is trying to signin
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

}

// SignUp create a user in the database
var SignUp = func(w http.ResponseWriter, r *http.Request) {

}
