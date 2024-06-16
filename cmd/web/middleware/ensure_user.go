package middleware

import (
	"context"
	"net/http"
	"sukasa/bookings/internal/users"
)

type contextKey string

const AuthUserID contextKey = "middleware.auth.user_id"

func EnsureUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]
		userService := users.GetUserService()
		user, ok := userService.VerifyToken(tokenString)
		if !ok {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AuthUserID, user.ID)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
