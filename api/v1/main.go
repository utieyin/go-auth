package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env '%v'", err)
	}
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST"),
	)
	port := os.Getenv("APP_PORT")
	a.Run(":8010" + port)
}
