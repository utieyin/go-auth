package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	facebookAccessToken = os.Getenv("SECRET_FACEBOOK_ACCESS_TOKEN")
)

//App DB
type App struct {
	Router *mux.Router
	DB     *sql.DB
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
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
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
func (a *App) Run(addr string) {}

// GetFacebookAccessToken returns the GH access token
func GetFacebookAccessToken() string {
	return facebookAccessToken
}
