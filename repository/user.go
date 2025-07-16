package repository

import (
	"afperdomo2/go/rest-ws/models"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetAll(ctx context.Context) ([]*models.User, error)
	GetById(ctx context.Context, id int64) (*models.User, error)
	Close() error // Método para cerrar la conexión a la base de datos
}

var implementation UserRepository

func SetUserRepository(repository UserRepository) {
	implementation = repository
}

func Close() error {
	if implementation != nil {
		return implementation.Close()
	}
	return nil
}

func CreateUser(ctx context.Context, user *models.User) error {
	return implementation.Create(ctx, user)
}

func GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return implementation.GetAll(ctx)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetById(ctx, id)
}
