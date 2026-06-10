package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/homecloud/go-api/internal/common"
	"github.com/homecloud/go-api/internal/jwt"
)

func AuthMiddleware(
	jwtService *jwt.Service,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.NewResponse(
				false,
				"Missing Authorization header",
				nil,
				nil,
			))
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.NewResponse(
				false,
				"Invalid or expired token",
				nil,
				nil,
			))
			return
		}
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
