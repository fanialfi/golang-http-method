package main

import (
	"golang-http-method/handler"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", handler.HandleIndex)

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
