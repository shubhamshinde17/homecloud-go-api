package auth

import "time"

type User struct {
	ID           int64
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
