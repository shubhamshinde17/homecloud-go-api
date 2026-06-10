package auth

import (
	"errors"

	"github.com/homecloud/go-api/internal/config"
	"github.com/homecloud/go-api/internal/jwt"
	"github.com/homecloud/go-api/internal/logger"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("EMAIL_ALREADY_EXISTS")
var ErrUserNotFound = errors.New("USER_NOT_FOUND")
var ErrInvalidCredentials = errors.New("INVALID_CREDENTIALS")
var JWTSecret = config.LoadConfig().JWTSecret

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

	logger.Error("Error Found: ", err)
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
	creationErr := s.repo.Create(user)
	logger.Error("Creation Error: ", creationErr)
	if creationErr != nil && !errors.Is(creationErr, ErrUserNotFound) {
		return nil, creationErr
	}
	return user, nil
}

func (s *Service) Login(req *LoginRequest) (*AccessTokenResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	accessToken, refreshToken, expiresIn, err := jwt.NewService(JWTSecret).GenerateTokens(user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	return &AccessTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

// func (s *Service) RefreshToken(req *RefreshTokenRequest) (*AccessTokenResponse, error) {
// 	claims, err := jwt.NewService(JWTSecret).ValidateRefreshToken(req.RefreshToken)
// 	if err != nil {
// 		return nil, err
// 	}
// 	accessToken, refreshToken, expiresIn, err := jwt.NewService(JWTSecret).GenerateTokens(claims.UserID, claims.Email)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &AccessTokenResponse{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 		ExpiresIn:    expiresIn,
// 	}, nil
// }

func (s *Service) GetUserByAccessToken(accessToken string) (*User, error) {
	claims, err := jwt.NewService(JWTSecret).ValidateToken(accessToken)
	if err != nil {
		logger.Error("Token validation error: ", err)
		return nil, ErrInvalidCredentials
	}
	user, err := s.repo.FindByID(claims.UserID)
	if err != nil {
		logger.Error("User not found: ", err)
		return nil, ErrUserNotFound
	}
	return user, nil
}

func MapUserToResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsActive:  user.IsActive,
	}
}
