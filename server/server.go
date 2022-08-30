package server

import (
	"log"
	"time"
	"github.com/joseluissanchez77/GoTesisApiArduino/routes"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/itsjamie/gin-cors"
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


	router.Run()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	// // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// // config.AllowAllOrigins = true
  
	// router.Use(cors.New(config))
	// router.Run()


	// router.Use(cors.Default())

	// // router := gin.Default()
	// corsConfig := cors.DefaultConfig()

	// corsConfig.AllowOrigins = []string{"*"}
	// // To be able to send tokens to the server.
	// corsConfig.AllowCredentials = true

	// // OPTIONS method for ReactJS
	// corsConfig.AddAllowMethods("POST, OPTIONS, GET, PUT, DELETE")

	// // Register the middleware
	// router.Use(cors.New(corsConfig))


// 	router = gin.Default()
//  // CORS for https://foo.com and https://github.com origins, allowing:
//  // - PUT and PATCH methods
//  // - Origin header
//  // - Credentials share
//  // - Preflight requests cached for 12 hours
//  router.Use(cors.New(cors.Config{
// //   AllowOrigins:     []string{"https://foo.com"},
//   AllowOrigins:     []string{"*"},
//   AllowMethods:     []string{ "POST, OPTIONS, GET, PUT, DELETE"},
//   AllowHeaders:     []string{"Origin"},
//   ExposeHeaders:    []string{"Content-Length"},
//   AllowCredentials: true,
//   AllowOriginFunc: func(origin string) bool {
//    return origin == "*"
//   },
//   MaxAge: 12 * time.Hour,
//  }))
//  router.Run()

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":"+s.port))
}