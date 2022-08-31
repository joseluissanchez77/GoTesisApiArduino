package main

import (
	"github.com/joseluissanchez77/GoTesisApiArduino/database"
	"github.com/joseluissanchez77/GoTesisApiArduino/server"
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