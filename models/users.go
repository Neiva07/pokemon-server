package models

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

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

	err := DB.GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please try again later."), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email address already in use."), false
	}

	return u.Message(false, "Requirement passed"), true

}

//Create function to create user account and hash the password into the database
func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	account.Password = string(hashedPassword)

	DB.GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create an account. Connection error.")
	}

	tk := Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	account.Token = tokenString
	account.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response

}

//Login function to handle with the request of the database
func Login(email string, password string) map[string]interface{} {

	account := &Account{}

	err := DB.GetDB().Table("accounts").Where("email = ?", email).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not register")
		}

		return u.Message(false, "Connection error. Please retry")
	}

	account.Password = "" // for safety

	tk := &Token{UserId: account.ID}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return u.Message(false, "something went wrong with the login system")
	}

	account.Token = tokenString

	resp := u.Message(true, "Logged in")
	resp["account"] = account
	return resp

}
