package main

import (
	"log"
	"net/http"
	"server/routers"
)

func main() {
	router := routers.SetupRouter()

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
