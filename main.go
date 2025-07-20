package main

import (
	"afperdomo2/go/rest-ws/handlers"
	"afperdomo2/go/rest-ws/middlewares"
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
	// Middleware de autenticación
	r.Use(middlewares.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	// Aquí puedes agregar más rutas según sea necesario
	r.HandleFunc("/signup", handlers.SingUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/user-info", handlers.GetUserFromTokenHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/posts/{id:[0-9]+}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/posts", handlers.CreatePostHandler(s)).Methods(http.MethodPost)
}
