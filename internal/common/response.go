package common

import "time"

func NewResponse(success bool, message string, data any, meta any) Response[any, any] {
	return Response[any, any]{
		Success:   success,
		Message:   message,
		Timestamp: time.Now(),
		Data:      &data,
		Meta:      &meta,
	}
}

func NewAccessTokenResponse(accessToken string, expiresIn int, refreshToken string) AccessTokenResponse {
	return AccessTokenResponse{
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
	}
}

func NewErrorResponse(message string) Response[any, any] {
	return NewResponse(
		false,
		message,
		nil,
		nil,
	)
}

type Response[T any, U any] struct {
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Data      *T        `json:"data,omitempty"`
	Meta      *U        `json:"meta,omitempty"`
}

type AccessTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}
