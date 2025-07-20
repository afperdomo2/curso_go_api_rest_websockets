package repository

import (
	"afperdomo2/go/rest-ws/models"
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	CreatePost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id int64) (*models.Post, error)

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

func GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return implementation.GetAllUsers(ctx)
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

func GetPostById(ctx context.Context, id int64) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}
