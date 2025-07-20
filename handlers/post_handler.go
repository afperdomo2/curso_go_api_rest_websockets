package handlers

import (
	"afperdomo2/go/rest-ws/models"
	"afperdomo2/go/rest-ws/repository"
	"afperdomo2/go/rest-ws/server"
	"afperdomo2/go/rest-ws/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UpsertPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostUpdateResponse struct {
	Message string `json:"message"`
}

func CreatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpsertPostRequest
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

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpsertPostRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		postIdStr := mux.Vars(r)["id"]
		if postIdStr == "" {
			http.Error(w, "Post ID is required", http.StatusBadRequest)
			return
		}

		postId, err := strconv.ParseInt(postIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid Post ID", http.StatusBadRequest)
			return
		}

		post := &models.Post{
			Id:      postId,
			Title:   req.Title,
			Content: req.Content,
		}

		err = repository.UpdatePost(r.Context(), postId, post)
		if err != nil {
			http.Error(w, "Error updating post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(PostUpdateResponse{
			Message: "Post updated successfully",
		})
	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIdStr := mux.Vars(r)["id"]
		if postIdStr == "" {
			http.Error(w, "Post ID is required", http.StatusBadRequest)
			return
		}

		// Convert postIdStr to int64
		postId, err := strconv.ParseInt(postIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid Post ID", http.StatusBadRequest)
			return
		}

		post, err := repository.GetPostById(r.Context(), postId)
		if err != nil {
			http.Error(w, "Error fetching post: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if post == nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}
}
