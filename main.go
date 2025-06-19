package main

import (
	"log"
	"net/http"

	"bakulos_grapghql/db"
)

func main() {
	db.ConnectDatabase()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
