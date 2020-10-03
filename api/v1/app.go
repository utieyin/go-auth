package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

// App sets up and runs the app
type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize sets up the db
func (a *App) Initialize(user, password, dbname, port, host string) {
	var err error
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	a.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("something went wrong")
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()

}
func (a *App) Run(addr string) {}
