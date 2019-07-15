package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"pokemon-server/models"
	u "pokemon-server/utils"
	"strings"
	"testing"
)

func TestEmptyTableAccounts(t *testing.T) {
	u.ClearTable(models.DB.GetDB(), "accounts")

	req, err := http.NewRequest("GET", "/accounts", nil)

	if err != nil {
		t.Error("unexpected error happend:", err)
	}

	response := u.ExecuteRequest(GetAllAccounts, req)

	u.CheckResponseCode(t, http.StatusOK, response.Code)

	expectedAnswer := `{"accounts":[],"message":"All accounts register","status":true}`
	if body := strings.TrimRight(response.Body.String(), "\n"); body != expectedAnswer {
		t.Errorf("Expected an empty array like %s. Got %s", expectedAnswer, body)
	}
}

func TestSignUp(t *testing.T) {

	type testUser struct {
		Acc      *models.Account
		response map[string]interface{}
	}

	tests := []testUser{
		testUser{
			Acc: &models.Account{
				Email:    "lucas@neiva.com",
				Password: "123467",
			},
			response: map[string]interface{}{"status": true, "message": "Account has been created"},
		},
		testUser{
			Acc: &models.Account{
				Email:    "lucas_neiva@a.com",
				Password: "ludhfuaodfh",
			},
			response: map[string]interface{}{"status": true, "message": "Account has been created"},
		},
		testUser{
			Acc: &models.Account{
				Email:    "lucasneiva.com",
				Password: "1243324",
			},
			response: map[string]interface{}{"status": false, "message": "Invalid Email, please try with a real one."},
		},
		testUser{
			Acc: &models.Account{
				Email:    "lucas@gaogh.com",
				Password: "121",
			},
			response: map[string]interface{}{"status": false, "message": "Password must be unless 6 character long"},
		},
	}
	payload := new(bytes.Buffer)

	for _, testCase := range tests {

		json.NewEncoder(payload).Encode(testCase.Acc)

		req, err := http.NewRequest("GET", "/api/signup", payload)

		if err != nil {
			t.Error("unexpected error happend:", err)
		}

		response := u.ExecuteRequest(SignUp, req)
		u.CheckResponseCode(t, http.StatusOK, response.Code)

		var body map[string]interface{}
		json.Unmarshal(response.Body.Bytes(), &body)

		if body["status"] != testCase.response["status"] {
			t.Errorf("Expected status '%v'. Got '%v'", testCase.response["status"], body["status"])
		}
		if body["message"] != testCase.response["message"] {
			t.Errorf("Expected message '%v'. Got '%v'", testCase.response["message"], body["message"])
		}
		if body["status"] == true {
			acc := body["account"].(map[string]interface{})

			if acc["email"] != (*testCase.Acc).Email {
				t.Errorf("Expected account '%v'. Got '%v'", (*testCase.Acc).Email, acc["email"])

			}
			//add more fields when restructure the Accounts
		}

	}

}

func TestGetAllAccounts(t *testing.T) {

	req, err := http.NewRequest("GET", "/accounts", nil)

	if err != nil {
		t.Error("unexpected error happend:", err)
	}

	response := u.ExecuteRequest(GetAllAccounts, req)

	u.CheckResponseCode(t, http.StatusOK, response.Code)

	accounts := &[]models.Account{}
	models.DB.GetDB().Table("accounts").Find(accounts)

}
