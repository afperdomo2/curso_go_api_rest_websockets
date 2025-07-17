package repository

import (
	"afperdomo2/go/rest-ws/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	Close() error // Método para cerrar la conexión a la base de datos
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func Close() error {
	if implementation != nil {
		return implementation.Close()
	}
	return nil
}

func CreateUser(ctx context.Context, user *models.User) error {
	return implementation.CreateUser(ctx, user)
}

func GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return implementation.GetAllUsers(ctx)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}
