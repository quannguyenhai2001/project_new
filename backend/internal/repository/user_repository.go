package repository

import (
	"backend/internal/model"
	"database/sql"
	"errors"
)

// UserRepository định nghĩa các phương thức tương tác với bảng users
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id int64) (*model.User, error)
}

// PostgresUserRepository là implementation của UserRepository với PostgreSQL
type PostgresUserRepository struct {
	db *sql.DB
}

// NewUserRepository tạo instance mới của PostgresUserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

// Create inserts a new user into the database
func (r *PostgresUserRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users (email, password, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// FindByEmail finds a user by their email
func (r *PostgresUserRepository) FindByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, password, first_name, last_name, created_at, updated_at
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	var user model.User
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// FindByID finds a user by their ID
func (r *PostgresUserRepository) FindByID(id int64) (*model.User, error) {
	query := `
		SELECT id, email, password, first_name, last_name, created_at, updated_at
		FROM users
		WHERE id = $1
		LIMIT 1
	`

	var user model.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
