package models

import "testing"

func TestCreate(t *testing.T) {

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
	for _, testCase := range tests {
		respGot := testCase.Acc.Create()
		for key, val := range respGot {
			if val != testCase.response[key] {
				t.Errorf("Expected %v, got %v at the response message", testCase.response, respGot)
			}
		}
	}

}
