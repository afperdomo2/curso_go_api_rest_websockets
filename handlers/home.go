// Package handlers contiene los manejadores HTTP para los diferentes endpoints de la API REST
// Cada handler se encarga de procesar las peticiones HTTP y generar las respuestas apropiadas
package handlers

import (
	"afperdomo2/go/rest-ws/server"
	"encoding/json"
	"net/http"
)

// HomeReponse representa la estructura de respuesta para el endpoint home
// Contiene un mensaje de bienvenida y el estado de la operación
type HomeReponse struct {
	Message string `json:"message"` // Mensaje de bienvenida de la API
	Status  bool   `json:"status"`  // Estado de la operación (true = exitoso)
}

// HomeHandler crea un manejador HTTP para el endpoint principal de la API
// Este endpoint sirve como punto de entrada y verificación de que la API está funcionando
//
// Parámetros:
//   - server: Instancia del servidor que proporciona acceso a la configuración
//
// Retorna:
//   - http.HandlerFunc: Función manejadora que procesa las peticiones HTTP
//
// Respuesta:
//   - Content-Type: application/json
//   - Status Code: 200 (OK)
//   - Body: JSON con mensaje de bienvenida y estado
func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Establece el tipo de contenido como JSON
		w.Header().Set("Content-Type", "application/json")
		// Establece el código de estado HTTP 200 (OK)
		w.WriteHeader(http.StatusOK)
		// Codifica y envía la respuesta JSON al cliente
		json.NewEncoder(w).Encode(HomeReponse{
			Message: "Welcome to the Go REST API",
			Status:  true,
		})
	}
}
