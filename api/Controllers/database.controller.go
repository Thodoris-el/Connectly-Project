package controllers

import (
	"flag"
	"log"
	"net/http"
	"os"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//struct for our server
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Db_user, Db_password, Db_host, Db_name, Db_port string) {
	var err error

	//dsn := Db_user + ":" + Db_password + "@tcp" + "(" + Db_host + ":" + Db_port + ")/" + Db_name + "?" + "parseTime=true&loc=Local"
	dsn := os.Getenv("DSN_NAME") + ":" + os.Getenv("DSN_PASSWORD") + "@tcp(aws-eu-west-2.connect.psdb.cloud)/connectly_project?tls=true"
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Cant connect to mysql database")
		log.Println(err)
	} else {
		log.Println("Connected to mysql database")
	}

	server.DB.Debug().AutoMigrate(&entity.Customer{}, &entity.Review{}, &entity.Conversation{}, &entity.Template{})
	server.Router = mux.NewRouter()
	server.startRoutes()
}

func (server *Server) Run() {
	//port := ":8000"
	port := flag.String("port", "8000", "specify a port")
	flag.Parse()

	log.Printf("Server started on %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, server.Router))
}
