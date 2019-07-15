package utils

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"
)

//CheckResponseCode is a function to test if the response code of http response is correct
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//ExecuteRequest is a function to handle with Requests inside the server
func ExecuteRequest(f func(w http.ResponseWriter, r *http.Request), req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(f)
	handler.ServeHTTP(rr, req)

	return rr
}

//ClearTable delete all rows from the a table to test propouse
func ClearTable(db *gorm.DB, table string) {
	deleteStr := fmt.Sprintf("DELETE FROM %v;", table)
	alterSeq := fmt.Sprintf("ALTER SEQUENCE %v_id_seq RESTART WITH 1;", table)
	db.Exec(deleteStr)
	db.Exec(alterSeq)

	log.Print("Table Cleared")
}
