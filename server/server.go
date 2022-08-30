package server

import (
	"log"
	"time"
	"github.com/joseluissanchez77/GoTesisApiArduino/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"os"
)

type Server struct{
	port string
	server *gin.Engine
}

func NewServer() Server{
	return Server{
		port: os.Getenv("PORT"),
		// port: "9001",
		server: gin.Default(),
	}
}

func (s *Server)Run(){
	router := routes.ConfigRoutes(s.server)

	router = gin.Default()
 // CORS for https://foo.com and https://github.com origins, allowing:
 // - PUT and PATCH methods
 // - Origin header
 // - Credentials share
 // - Preflight requests cached for 12 hours
 router.Use(cors.New(cors.Config{
//   AllowOrigins:     []string{"https://foo.com"},
  AllowOrigins:     []string{"*"},
  AllowMethods:     []string{ "POST, OPTIONS, GET, PUT, DELETE"},
  AllowHeaders:     []string{"Origin"},
  ExposeHeaders:    []string{"Content-Length"},
  AllowCredentials: true,
  AllowOriginFunc: func(origin string) bool {
   return origin == "*"
  },
  MaxAge: 12 * time.Hour,
 }))
 router.Run()

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":"+s.port))
}