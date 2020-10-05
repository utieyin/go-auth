package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/utieyin/go-auth/api/v1/config"
)

func main() {
	var err error
	err = godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env '%v'", err)
	}
	a := config.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	port := os.Getenv("APP_PORT")
	a.Run(":8010" + port)
}
