package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/homecloud/go-api/internal/common"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Method not Allowed",
		))
		return
	}

	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Invalid request payload",
		))
		return
	}

	user, err := h.service.Signup(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if errors.Is(err, ErrEmailAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(common.NewErrorResponse(
				"Email already exists",
			))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Internal Server Error",
		))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	mappedUser := MapUserToResponse(user)
	json.NewEncoder(w).Encode(common.NewResponse(
		true,
		"User created successfully",
		mappedUser,
		nil,
	))
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Method not Allowed",
		))
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Invalid request payload",
		))
		return
	}

	loginResponse, err := h.service.Login(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case errors.Is(err, ErrUserNotFound):
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(common.NewErrorResponse(
				"User not found",
			))
			return
		case errors.Is(err, ErrInvalidCredentials):
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.NewErrorResponse(
				"Invalid email or password",
			))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(common.NewErrorResponse(
				"Internal Server Error",
			))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(common.NewResponse(
		true,
		"Login successful",
		loginResponse,
		nil,
	))
}

// func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		json.NewEncoder(w).Encode(common.NewResponse(
// 			false,
// 			"Method not Allowed",
// 			nil,
// 			nil,
// 		))
// 		return
// 	}
// 	err := h.service.Logout(r)
// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(common.NewResponse(
// 			false,
// 			"Internal Server Error",
// 			nil,
// 			nil,
// 		))
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(common.NewResponse(
// 		true,
// 		"Logout successful",
// 		nil,
// 		nil,
// 	))
// }

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Method not Allowed",
		))
		return
	}
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Missing access token",
		))
		return
	}
	user, err := h.service.GetUserByAccessToken(accessToken)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if errors.Is(err, ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(common.NewErrorResponse(
				"User not found",
			))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.NewErrorResponse(
			"Internal Server Error",
		))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mappedUser := MapUserToResponse(user)
	json.NewEncoder(w).Encode(common.NewResponse(
		true,
		"User retrieved successfully",
		mappedUser,
		nil,
	))
}
