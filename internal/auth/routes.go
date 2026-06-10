package auth

import (
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	handler := NewHandler(NewService(NewRepository(db)))
	handler.RegisterRoutes(mux, handler)
	return mux
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux, authHandler *Handler) {
	mux.HandleFunc("/signup", authHandler.SignupHandler)
	mux.HandleFunc("/login", authHandler.LoginHandler)
	mux.HandleFunc("/me", authHandler.GetUserHandler)
}
