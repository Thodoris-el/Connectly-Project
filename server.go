package main

import (
	"log"
	"os"

	controllers "github.com/Thodoris-el/Coonectly-Project/api/Controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Initialize request

func main() {
	//Load env
	err := godotenv.Load()
	if err != nil {
		log.Println("error while loading .env file")
	} else {
		log.Println("We are getting the .env values")
	}

	Db_user := os.Getenv("DB_USERNAME")
	Db_password := os.Getenv("DB_PASSWORD")
	Db_host := os.Getenv("DB_HOST")
	Db_port := os.Getenv("DB_PORT")
	Db_name := os.Getenv("DB_NAME")

	server.Initialize(Db_user, Db_password, Db_host, Db_name, Db_port)
	server.Run()
}
