package main

import (
	// "net/http"
	// "log"
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