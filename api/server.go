package api

import (
	"log"
	"os"

	controllers "github.com/Thodoris-el/Coonectly-Project/api/Controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	//load values from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("error while loading env ", err)
	}
}

//Run api server
func Run() {
	//load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("error while loading .env file")
	} else {
		log.Println("We are getting the .env values")
	}

	//Get db credentials from .env
	Db_user := os.Getenv("DB_USERNAME")
	Db_password := os.Getenv("DB_PASSWORD")
	Db_host := os.Getenv("DB_HOST")
	Db_port := os.Getenv("DB_PORT")
	Db_name := os.Getenv("DB_NAME")
	Dsn_Name := os.Getenv("DSN_NAME")
	Dsn_Password := os.Getenv("DSN_PASSWORD")

	//Initialize DB -> connect and migrate
	server.Initialize(Db_user, Db_password, Db_host, Db_name, Db_port, Dsn_Name, Dsn_Password)
	//run server
	server.Run()
}
