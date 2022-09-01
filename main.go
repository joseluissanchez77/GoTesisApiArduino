package main

import (
	// "time"

    // "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"github.com/joseluissanchez77/GoTesisApiArduino/server"
	// "github.com/gin-contrib/cors"
	// "time"
)

func main(){

	database.StartDB()
	
	server := server.NewServer()
//  router = gin.New() 
// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"*"},
// 		AllowMethods:     []string{"PUT", "PATCH", "GET"},
// 		AllowHeaders:     []string{"Origin"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		AllowOriginFunc: func(origin string) bool {
// 			return origin == "https://github.com"
// 		},
// 		MaxAge: 12 * time.Hour,
// 	}))
	// router := gin.Default()

	// router.Use(cors.New(cors.Config{
    //     AllowOrigins:     []string{"*"},
    //     AllowMethods:     []string{"PUT", "PATCH","GET","POST"},
    //     AllowHeaders:     []string{"Origin"},
    //     ExposeHeaders:    []string{"Content-Length"},
    //     AllowCredentials: true,
    //     AllowOriginFunc: func(origin string) bool {
    //         return origin == "https://github.com"
    //     },
    //     MaxAge: 12 * time.Hour,
    // }))
	server.Run()
}

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello..., 世界")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":9001", nil))
// }