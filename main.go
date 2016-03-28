package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Printf("Listening...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
