package main

import (
	"log"
	"net/http"

	"app/controllers"
)

func main() {

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/echo", controllers.WebSocket)

	log.Fatal(http.ListenAndServe(":8080", nil))
}