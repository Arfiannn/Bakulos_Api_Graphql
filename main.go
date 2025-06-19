package main

import (
	"bakulos_grapghql/auth"
	"bakulos_grapghql/db"
	"bakulos_grapghql/routes/schema"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	db.ConnectDatabase()

	// Buat GraphQL schema (Query + Mutation)
	schema := schema.NewSchema()

	http.HandleFunc("/login", auth.LoginHandler)

	// Setup handler dengan GraphiQL UI
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", auth.AuthMiddleware(h))

	log.Println("🚀 Server GraphQL berjalan di: http://localhost:8080/graphql")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
