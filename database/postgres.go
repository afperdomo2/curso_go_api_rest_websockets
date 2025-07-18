package database

import (
	"afperdomo2/go/rest-ws/models"
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Close() error {
	// Cierra la conexión a la base de datos
	if r.db != nil {
		return r.db.Close()
	}
	return nil // Si no hay conexión, retorna nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user *models.User) error {
	// Utiliza un contexto para manejar la operación de forma segura
	// Realiza una inserción en la base de datos para crear un nuevo usuario
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (r *PostgresRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	// Realiza una consulta a la base de datos para obtener todos los usuarios
	// Utiliza un contexto para manejar la operación de forma segura
	rows, err := r.db.QueryContext(ctx, "SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Asegura que los recursos se liberen después de la consulta

	var users []*models.User // Inicializa un slice para almacenar los usuarios

	// Itera sobre los resultados de la consulta y los agrega al slice
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *PostgresRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	// Realiza una consulta a la base de datos para encontrar un usuario por su ID
	// Utiliza un contexto para manejar la operación de forma segura
	row := r.db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	var user models.User
	// Escanea los resultados de la consulta en la estructura del usuario
	if err := row.Scan(&user.Id, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err // Error al escanear los resultados
	}
	return &user, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	var user models.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
