package repositories

import (
	"afperdomo2/go/rest-ws/models"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindAll(ctx context.Context) ([]*models.User, error)
	FindById(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int64) error
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

func Create(ctx context.Context, user *models.User) error {
	return implementation.Create(ctx, user)
}

func FindAll(ctx context.Context) ([]*models.User, error) {
	return implementation.FindAll(ctx)
}

func FindById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.FindById(ctx, id)
}

func Update(ctx context.Context, user *models.User) error {
	return implementation.Update(ctx, user)
}

func Delete(ctx context.Context, id int64) error {
	return implementation.Delete(ctx, id)
}
