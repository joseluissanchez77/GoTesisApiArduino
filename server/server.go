package server

import (
	"log"
	// "time"
	"github.com/joseluissanchez77/GoTesisApiArduino/routes"
	"github.com/gin-gonic/gin"
	"os"
	// "github.com/itsjamie/gin-cors"
	"github.com/gin-contrib/cors"
	
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

	// router = gin.New()  
	// router.Use(cors.Default())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":"+s.port))
}



func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}