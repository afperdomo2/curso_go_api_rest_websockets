package repository

import (
	"afperdomo2/go/rest-ws/models"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	CreatePost(ctx context.Context, post *models.Post) error
	UpdatePost(ctx context.Context, id int64, changes *models.Post) error
	GetPostById(ctx context.Context, id int64) (*models.Post, error)
	DeletePost(ctx context.Context, id int64, userId int64) error
	GetAllPosts(ctx context.Context, page int64, limit int64) ([]*models.Post, error)

	Close() error // Método para cerrar la conexión a la base de datos
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func Close() error {
	if implementation != nil {
		return implementation.Close()
	}
	return nil
}

// User
func CreateUser(ctx context.Context, user *models.User) error {
	return implementation.CreateUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

// Post
func CreatePost(ctx context.Context, post *models.Post) error {
	return implementation.CreatePost(ctx, post)
}

func UpdatePost(ctx context.Context, id int64, changes *models.Post) error {
	return implementation.UpdatePost(ctx, id, changes)
}

func GetPostById(ctx context.Context, id int64) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func DeletePost(ctx context.Context, id int64, userId int64) error {
	return implementation.DeletePost(ctx, id, userId)
}

func GetAllPosts(ctx context.Context, page int64, limit int64) ([]*models.Post, error) {
	return implementation.GetAllPosts(ctx, page, limit)
}
