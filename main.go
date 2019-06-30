package main

import (
	"encoding/json"
	"log"
	"net/http"
	controllers "pokemon-server/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", mainFunc).Methods("GET")
	router.HandleFunc("/api/user/signin", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/signup", controllers.SignUp).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func mainFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	x := struct {
		Data string
	}{
		Data: "Matheus Otário",
	}
	err := json.NewEncoder(w).Encode(x)
	log.Println(x)

	if err != nil {
		log.Fatal(err)
	}
	return
}
