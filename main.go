package main

import (
	"os"
	"time"
	"github.com/joseluissanchez77/GoTesisApiArduino/controllers"
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)



func main(){


	router := gin.Default()
    // router.Use(cors.Default())
    // router.GET("/", func(c *gin.Context) {
    //     c.JSON(200, gin.H{"message": "CORS works!"})
    // })
    // router.Run(":9001")

	
	database.StartDB()
	
	// server := server.NewServer()
	// server.Run()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH","GET","POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	  }))

	router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Api Arduino!"})
    })
	router.GET("/api/v1/catalog-data", controllers.ShowAllCatalogData)
	router.Run(os.Getenv("PORT"))	
	// router.Run(":9001")	
		
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