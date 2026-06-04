package auth

import (
	"database/sql"
	"errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	row := r.db.QueryRow(`SELECT id, first_name, last_name, email, password_hash, is_active FROM users WHERE email = $1`, email)
	user := &User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) FindByID(id int64) (*User, error) {
	row := r.db.QueryRow(`SELECT id, first_name, last_name, email, password_hash, is_active FROM users WHERE id = $1`, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) Create(user *User) error {
	return r.db.QueryRow(`INSERT INTO users (first_name, last_name, email, password_hash, is_active) VALUES ($1, $2, $3, $4, $5)`,
		user.FirstName, user.LastName, user.Email, user.PasswordHash, user.IsActive).Scan(&user.ID)
}
