package models

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {

	t.Parallel()
	type testUser struct {
		Acc      *Account
		response map[string]interface{}
	}

	tests := []testUser{
		testUser{
			Acc: &Account{
				Email:    "lucas@neiva.com",
				Password: "123467",
			},
			response: map[string]interface{}{"status": true, "message": "Account has been created"},
		},
		testUser{
			Acc: &Account{
				Email:    "lucas_neiva@a.com",
				Password: "ludhfuaodfh",
			},
			response: map[string]interface{}{"status": true, "message": "Account has been created"},
		},
		testUser{
			Acc: &Account{
				Email:    "lucasneiva.com",
				Password: "1243324",
			},
			response: map[string]interface{}{"status": false, "message": "Invalid Email, please try with a real one."},
		},
		testUser{
			Acc: &Account{
				Email:    "lucas@gaogh.com",
				Password: "121",
			},
			response: map[string]interface{}{"status": false, "message": "Password must be unless 6 character long"},
		},
	}
	for ind, testCase := range tests {
		fmt.Println(ind)
		respGot := testCase.Acc.Create()
		if respGot["status"] != testCase.response["status"] || respGot["message"] != testCase.response["message"] {
			t.Errorf("Expected %v, got %v at the response message. Case %v", testCase.response, respGot, (*testCase.Acc).Email)
		}
	}

}
