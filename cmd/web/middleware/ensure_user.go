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
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userService := users.GetUserService()
		user, ok := userService.GetCurrentUser(username, password)
		if !ok {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AuthUserID, user.ID)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
