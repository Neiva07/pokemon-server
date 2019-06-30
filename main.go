package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	controllers "pokemon-server/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", mainFunc).Methods("GET")
	router.HandleFunc("/api/user/signin", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/signup", controllers.SignUp).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
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
