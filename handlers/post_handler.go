package handlers

import (
	"afperdomo2/go/rest-ws/models"
	"afperdomo2/go/rest-ws/repository"
	"afperdomo2/go/rest-ws/server"
	"afperdomo2/go/rest-ws/services"
	"encoding/json"
	"net/http"
)

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreatePostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user := services.UserServiceInstance.GetUserFromToken(r, s, w)

		post := models.Post{
			Title:   req.Title,
			Content: req.Content,
			UserID:  user.Id,
		}

		err := repository.CreatePost(r.Context(), &post)
		if err != nil {
			http.Error(w, "Error creating post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	}
}
