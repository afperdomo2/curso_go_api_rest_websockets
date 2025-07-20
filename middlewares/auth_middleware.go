package middlewares

import (
	"afperdomo2/go/rest-ws/server"
	"net/http"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	PUBLIC_ENDPOINTS = []string{
		"/signup",
		"/login",
	}
)

// Si la ruta no está en la lista de endpoints públicos, se requiere verificación
func shouldCheckToken(route string) bool {
	return !slices.Contains(PUBLIC_ENDPOINTS, route)
}

// CheckAuthMiddleware verifica el token JWT en las rutas protegidas
// Si el token es válido, permite el acceso a la ruta; de lo contrario, retorna un error 401 Unauthorized
// Este middleware se aplica a todas las rutas excepto a las que están en PUBLIC_ENDPOINTS
func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Si la ruta no requiere verificación de token, se pasa directamente al siguiente handler
			if !shouldCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			// Si la ruta requiere verificación de token, se verifica el JWT
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (any, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Si el token es válido, se pasa al siguiente handler
			next.ServeHTTP(w, r)
		})
	}
}
