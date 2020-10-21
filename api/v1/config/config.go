package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/utieyin/go-auth/api/v1/utils"

	_ "github.com/lib/pq"
)

//App DB
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/tokens", a.getProduct).Methods("GET")
}

// Initialize an app
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("something went wrong")
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	value, err := utils.DebugToken(r)
	if err != nil {
		log.Fatal("Token is invalid")
	}
	fmt.Println("valid token: ", value)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Run This is to run the app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))

}

// GetFacebookAccessToken returns the GH access token
// func GetFacebookAccessToken() string {
// 	return facebookAccessToken
// }
