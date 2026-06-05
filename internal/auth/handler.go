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
		json.NewEncoder(w).Encode(common.NewResponse(
			false,
			"Method not Allowed",
			nil,
			nil,
		))
		return
	}

	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.NewResponse(
			false,
			"Invalid request payload",
			nil,
			nil,
		))
		return
	}

	user, err := h.service.Signup(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if errors.Is(err, ErrEmailAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(common.NewResponse(
				false,
				"Email already exists",
				nil,
				nil,
			))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.NewResponse(
			false,
			"Internal Server Error",
			nil,
			nil,
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
