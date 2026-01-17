package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/amdaaraijen/Learn-API/internal/authctx"
	"github.com/amdaaraijen/Learn-API/internal/pkg/token"
)

func JWTAuth(jwtMaker *token.JWTMaker) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			claims, err := jwtMaker.VerifyToken(parts[1])

			if err != nil {
				http.Error(w, "invalid or expired authorization token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), authctx.UserIDKey, claims.UserID)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
