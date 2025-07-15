package main

import (
	"afperdomo2/go/rest-ws/handlers"
	"afperdomo2/go/rest-ws/server"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, error := server.NewServer(context.Background(), &server.ServerConfig{
		Port:        ":" + PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})
	if error != nil {
		log.Fatalf("Error creating server: %v", error)
	}
	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	// Aquí puedes agregar más rutas según sea necesario
}
