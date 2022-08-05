package server

import (
	"log"
	"github.com/joseluissanchez77/GoTesisApiArduino/routes"
	"github.com/gin-gonic/gin"
	"os"
)

type Server struct{
	port string
	server *gin.Engine
}

func NewServer() Server{
	return Server{
		port: os.Getenv("PORT"),
		server: gin.Default(),
	}
}

func (s *Server)Run(){
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":"+s.port))
}