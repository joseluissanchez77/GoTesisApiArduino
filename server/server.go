package server

import (
	// "net/http"
	"log"
	// "time"
	"github.com/joseluissanchez77/GoTesisApiArduino/routes"
	"github.com/gin-gonic/gin"
	// "os"

)

type Server struct{
	port string
	server *gin.Engine
}

func NewServer() Server{
	return Server{
		// port: os.Getenv("PORT"),
		port: "9001",
		server: gin.Default(),
	}
}

func (s *Server)Run(){

	router := routes.ConfigRoutes(s.server)


	// router.Use(Cors())

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":"+s.port))
}


// func Cors() gin.HandlerFunc {
//     return func(c *gin.Context) {
        
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") 
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") 
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers") 
// 		c.Writer.Header().Set("Access-Control-Max-Age", "172800") 
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")                                                                                                                                                                                                                          
        
//         c.Next()
//     }
// }


