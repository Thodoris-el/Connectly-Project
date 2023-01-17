package controllers

import (
	"log"
	"net/http"

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
	DSN := "jegaa58mmfhaikonpe0y:pscale_pw_ObRIK1EcJpcgtA3P47p6qShtWH5Xr1fFM108z46JuJY@tcp(aws-eu-west-2.connect.psdb.cloud)/connectly_project?tls=true"
	server.DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

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
	port := ":8000"
	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}
