package api

import (
	"context"
	"log"
	"net/http"
	"strings"
	"ticket/internal/auth"
)

type contextKey string

const UserContextKey = contextKey("user")

func (app *Application) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("No Authorization header found")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid Authorization header format")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		log.Printf("Extracted token: %s", tokenString) // For debugging

		claims, err := app.Authenticator.ValidateToken(tokenString)
		if err != nil {
			log.Println("Invalid token")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func (app *Application) GetUserFromContext(ctx context.Context) *auth.Claims {
	user, _ := ctx.Value(UserContextKey).(*auth.Claims)
	return user
}
