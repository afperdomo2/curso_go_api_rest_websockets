package handlers

import (
	"afperdomo2/go/rest-ws/models"
	"afperdomo2/go/rest-ws/repository"
	"afperdomo2/go/rest-ws/server"
	"encoding/json"
	"net/http"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func SingUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signupRequest SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&signupRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		var newUser = models.User{
			Email:    signupRequest.Email,
			Username: signupRequest.Username,
			Password: signupRequest.Password,
		}
		err := repository.CreateUser(r.Context(), &newUser)

		if err != nil {
			http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SignupResponse{
			Id:       newUser.Id,
			Email:    newUser.Email,
			Username: newUser.Username,
		})
	}
}
