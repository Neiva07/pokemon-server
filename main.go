package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	application "pokemon-server/app"
	controllers "pokemon-server/controllers"

	"github.com/gorilla/mux"
)

//App instance to build the app
type App struct {
	Router *mux.Router
}

//ExecuteRequest is a method to execute http requests

//UseRouter Method to initilize all routers of the app
func (app *App) UseRouter() {
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/", mainFunc).Methods("GET")
	app.Router.HandleFunc("/api/users/signin", controllers.Authenticate).Methods("POST")
	app.Router.HandleFunc("/api/users/signup", controllers.SignUp).Methods("POST")
	app.Router.HandleFunc("/api/users", controllers.GetAllAccounts).Methods("GET")

	app.Router.Use(application.JwtAuthentication)

}

func main() {

	app := App{}

	app.UseRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
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
