package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64  `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type Service struct {
	secret string
}

func NewService(secret string) *Service {
	return &Service{
		secret: secret,
	}
}

func (s *Service) GenerateAccessToken(
	userID int64,
	email string,
) (string, error) {

	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(15 * time.Minute),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(s.secret),
	)
}

func (s *Service) GenerateRefreshToken(
	userID int64,
	email string,
) (string, error) {

	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					7 * 24 * time.Hour,
				),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(s.secret),
	)
}

func (s *Service) ValidateToken(
	tokenString string,
) (*Claims, error) {
	tokenString, err := getTokenFromHeader(tokenString)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func (s *Service) GenerateTokens(
	userID int64,
	email string,
) (string, string, int64, error) {
	accessToken, err := s.GenerateAccessToken(userID, email)
	if err != nil {
		return "", "", 0, err
	}
	refreshToken, err := s.GenerateRefreshToken(userID, email)
	if err != nil {
		return "", "", 0, err
	}
	return accessToken, refreshToken, 15 * 60, nil
}

func getTokenFromHeader(authHeader string) (string, error) {
	tokenString := strings.TrimSpace(authHeader)
	if tokenString == "" {
		return "", jwt.ErrTokenMalformed
	}
	if strings.Contains(tokenString, " ") {
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) == 2 {
			tokenString = strings.TrimSpace(parts[1])
		}
	}
	return tokenString, nil
}
