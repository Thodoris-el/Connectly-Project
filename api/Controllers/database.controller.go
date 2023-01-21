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

//Initialize Database
func (server *Server) Initialize(Db_user, Db_password, Db_host, Db_name, Db_port, Dsn_Name, Dsn_Password string) {

	var err error

	//use for a local database -> uncomment the below line
	//dsn := Db_user + ":" + Db_password + "@tcp" + "(" + Db_host + ":" + Db_port + ")/" + Db_name + "?" + "parseTime=true&loc=Local"
	//use for a hosted database -> uncomment the below line ( host used: https://planetscale.com/)
	dsn := Dsn_Name + ":" + Dsn_Password + "@tcp(aws-eu-west-2.connect.psdb.cloud)/connectly_project?tls=true&parseTime=true&loc=Local"
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
	port := ":5000"
	//port := flag.String("port", "8000", "specify a port")
	//flag.Parse()

	//serverWithTimeOut := http.TimeoutHandler(server.Router, time.Second*10, "Timeout!")
	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}
