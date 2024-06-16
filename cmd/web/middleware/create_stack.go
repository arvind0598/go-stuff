package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// CreateStack creates a middleware stack.
// Its just easier to reason about compared to a nested function call.
func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}
