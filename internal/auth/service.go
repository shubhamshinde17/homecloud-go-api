package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("EMAIL_ALREADY_EXISTS")
var ErrUserNotFound = errors.New("USER_NOT_FOUND")

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *Service) Signup(req *SignupRequest) (*User, error) {
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	passwordHash, err := HashPassword(req.Password)

	user := &User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: passwordHash,
		IsActive:     true,
	}
	err = s.repo.Create(user)
	if err != nil && !errors.Is(err, ErrUserNotFound) {
		return nil, err
	}
	return user, nil
}
