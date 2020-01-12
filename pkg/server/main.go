package server

import (
	"log"

	"github.com/danymarita/go-gql-server/internal/handlers"
	"github.com/gin-gonic/gin"
)

var HOST, PORT string

func init() {
	HOST = "localhost"
	PORT = "8080"
}

func Run() {
	r := gin.Default()
	// Setup a route
	r.GET("/ping", handlers.Ping())

	log.Printf("Running at http://%s:%s", HOST, PORT)
	log.Fatalln(r.Run(HOST + ":" + PORT))
}
