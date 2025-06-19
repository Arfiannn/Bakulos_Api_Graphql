package main

import (
	"log"
	"net/http"

	"bakulos_grapghql/auth"
	"bakulos_grapghql/db"
	"bakulos_grapghql/routes/schema"
	"bakulos_grapghql/routes/websocket"

	"github.com/graphql-go/handler"
)

func main() {
	db.ConnectDatabase()

	// Buat GraphQL schema (Query + Mutation)
	schema := schema.NewSchema()

	// Route WebSocket
	http.HandleFunc("/ws", websocket.HandleWebSocket)
	http.HandleFunc("/login", auth.LoginHandler)

	// Setup handler dengan GraphiQL UI
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", auth.AuthMiddleware(h))

	log.Println("🚀 Server GraphQL berjalan di: http://localhost:8080/graphql")
	log.Println("🚀 WebSocket berjalan di: ws://192.168.1.9:8080/ws")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
