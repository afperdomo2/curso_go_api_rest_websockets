package utils

import (
	"afperdomo2/go/rest-ws/models"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	ErrMissingAuthHeader = errors.New("missing Authorization header")
	ErrInvalidToken      = errors.New("invalid token")
	ErrInvalidClaims     = errors.New("invalid token claims")
)

// ExtractTokenFromRequest extrae el token JWT del header Authorization de la request
// Esta es una función utilitaria pura que solo se encarga de la extracción
func ExtractTokenFromRequest(r *http.Request) (string, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
	if tokenString == "" {
		return "", ErrMissingAuthHeader
	}
	return tokenString, nil
}

// ParseAndValidateToken parsea y valida un token JWT, devolviendo los claims
// Esta función es pura y no tiene efectos secundarios
func ParseAndValidateToken(tokenString string, jwtSecret string) (*models.AppClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidClaims
}
