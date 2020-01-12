package server

import (
	"log"

	"github.com/danymarita/go-gql-server/internal/handlers"
	"github.com/danymarita/go-gql-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

var host, port string

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
}

func Run() {
	r := gin.Default()
	// Setup a route
	r.GET("/ping", handlers.Ping())

	log.Printf("Running at http://%s:%s", host, port)
	log.Fatalln(r.Run(host + ":" + port))
}
