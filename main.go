package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, 世界")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9001", nil))
}