package middleware

import (
	"context"
	"net/http"
	"sukasa/bookings/internal/users"
)

type contextKey string

const AuthUserID contextKey = "middleware.auth.user_id"

// Sets the expectation going forward that a user exists and is authenticated.
// This is supposed to be a parallel to how we set up instance variables in RoR.
// I couldn't find a way to do that in Go, so I'm using context instead.
// The logic is that a search by user ID is probably faster than a search by email/password.
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
