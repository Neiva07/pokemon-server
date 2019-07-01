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

// validate function to validate the login
func (account *Account) validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {

		return u.Message(false, "Invalid Email, please try with a real one."), false
	}
	if len(account.Password) < 6 {
		return u.Message(false, "Password must be unless 6 character long"), false
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please try again later."), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use."), false
	}

	return u.Message(false, "Requirement passed"), true

}

// func (account *Account) Create() map[string]interface{} {

// 	if resp, ok := account.validate(); !ok {
// 		return resp
// 	}

// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

// 	account.Password = string(hashedPassword)

// 	GetDB().Create(account)

// 	if account.ID <= 0 {
// 		return u.Message(false, "Failed to create an account. Connection error.")
// 	}

// }
