package models

import "testing"

func TestCreate(t *testing.T) {

	type testUser struct {
		Account
		response map[string]interface{}
	}

	tests := []testUser{
		testUser{
			Account: Account{
				Email:    "lucas@neiva.com",
				Password: "123467",
			},
			response: map[string]interface{}{"status": true, "message": "Account has been created"},
		},
		testUser{
			Account: Account{
				Email:    "lucas_neiva@a.com",
				Password: "ludhfuaodfh",
			},
			response: map[string]interface{}{"status": true, "message": "Account has been created"},
		},
		testUser{
			Account: Account{
				Email:    "lucasneiva.com",
				Password: "1243324",
			},
			response: map[string]interface{}{"status": false, "message": "Invalid Email, please try with a real one."},
		},
		testUser{
			Account: Account{
				Email:    "lucas@gaogh.com",
				Password: "121",
			},
			response: map[string]interface{}{"status": false, "message": "Password must be unless 6 character long"},
		},
	}

	for acc, resp := range tests {

	}

}
