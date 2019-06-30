package models

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	u "pokemon-server/utils"
)

// Token to represent the token from the users
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

func (account *Account) validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {

		return u.Message(false, "Invalid Email, please try with a real one."), false
	}
	if len(account.Password) < 6 {
		return u.Message(false, "Password must be unless 6 character long"), false
	}

}
