// Package server proporciona la infraestructura básica para un servidor HTTP REST
// Incluye configuración, inicialización y gestión del ciclo de vida del servidor
package server

import (
	"afperdomo2/go/rest-ws/database"
	"afperdomo2/go/rest-ws/repository"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server define la interfaz que debe implementar cualquier servidor
// Proporciona acceso a la configuración del servidor
type Server interface {
	Config() *ServerConfig
}

// ServerConfig contiene todos los parámetros de configuración necesarios para el servidor
type ServerConfig struct {
	Port        string // Puerto en el que el servidor escuchará (ej: ":8080")
	JWTSecret   string // Clave secreta para firmar y verificar tokens JWT
	DatabaseURL string // URL de conexión a la base de datos
}

// Broker es la implementación concreta del servidor HTTP
// Encapsula la configuración y el router de rutas
type Broker struct {
	config *ServerConfig // Configuración del servidor
	router *mux.Router   // Router HTTP para manejar las rutas
}

// Config devuelve la configuración actual del broker
// Implementa la interfaz Server
func (b *Broker) Config() *ServerConfig {
	return b.config
}

// NewServer crea una nueva instancia del servidor HTTP
// Valida que todos los parámetros de configuración requeridos estén presentes
// Retorna un error si algún parámetro obligatorio está vacío
//
// Parámetros:
//   - ctx: Contexto para la creación del servidor
//   - config: Configuración del servidor que incluye puerto, JWT secret y URL de base de datos
//
// Retorna:
//   - *Broker: Instancia del servidor configurada
//   - error: Error si la configuración es inválida
func NewServer(ctx context.Context, config *ServerConfig) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port must be specified")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("JWT secret must be specified")
	}
	if config.DatabaseURL == "" {
		return nil, errors.New("database URL must be specified")
	}
	// Crea una nueva instancia del broker con la configuración y un router vacío
	// El router se inicializa aquí para que esté listo para usar al iniciar el servidor
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

// Start inicia el servidor HTTP y lo pone en modo de escucha
// Configura las rutas usando la función binder proporcionada y luego
// inicia el servidor en el puerto especificado en la configuración
//
// Parámetros:
//   - binder: Función que recibe el servidor y router para configurar las rutas
//
// Nota: Esta función bloquea la ejecución hasta que el servidor se detenga
// En caso de error al iniciar el servidor, el programa terminará con log.Fatal
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	repo, err := database.NewPostgresRepository(b.config.DatabaseURL)
	if err != nil {
		log.Fatal("❌ Error connecting to database:", err)
	}
	repository.SetRepository(repo)

	log.Println("🚀 Server started on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("❌ Error starting server:", err)
	}
}
