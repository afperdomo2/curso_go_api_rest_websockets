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

func (r *PostgresRepository) CreatePost(ctx context.Context, post *models.Post) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO posts (title, content, user_id) VALUES ($1, $2, $3)", post.Title, post.Content, post.UserID)
	return err
}

func (r *PostgresRepository) GetPostById(ctx context.Context, id int64) (*models.Post, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, title, content, user_id FROM posts WHERE id = $1", id)

	var post models.Post
	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostgresRepository) UpdatePost(ctx context.Context, id int64, changes *models.Post) error {
	_, err := r.db.ExecContext(ctx, "UPDATE posts SET title = $1, content = $2 WHERE id = $3", changes.Title, changes.Content, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) DeletePost(ctx context.Context, id int64, userId int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 AND user_id = $2", id, userId)
	return err
}

func (r *PostgresRepository) GetAllPosts(ctx context.Context, page int64, limit int64) ([]*models.Post, error) {
	offset := (page - 1) * limit
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, content, user_id, created_at FROM posts ORDER BY id DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.UserID, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
