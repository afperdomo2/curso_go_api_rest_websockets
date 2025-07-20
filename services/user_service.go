package services

import (
	"afperdomo2/go/rest-ws/models"
	"afperdomo2/go/rest-ws/repository"
	"afperdomo2/go/rest-ws/server"
	"afperdomo2/go/rest-ws/utils"
	"context"
	"errors"
	"net/http"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// UserService contiene la lógica de negocio relacionada con usuarios
type UserService struct{}

// GetUserFromToken combina las utilidades de JWT con la lógica de negocio
// para obtener un usuario completo desde un token en la request
func (us *UserService) GetUserFromToken(r *http.Request, s server.Server, w http.ResponseWriter) *models.User {
	// Usar utilidad para extraer token
	tokenString, err := utils.ExtractTokenFromRequest(r)
	if err != nil {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return nil
	}

	// Usar utilidad para validar token y obtener claims
	claims, err := utils.ParseAndValidateToken(tokenString, s.Config().JWTSecret)
	if err != nil {
		http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
		return nil
	}

	// Lógica de negocio: obtener usuario de la base de datos
	user, err := repository.GetUserById(context.Background(), claims.UserId)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return nil
	}

	return user
}

// Instancia global del servicio (patrón Singleton simple)
var UserServiceInstance = &UserService{}
