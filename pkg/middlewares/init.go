package middlewares

import "ticket/internal/modules/auth/services"

// Init initializes the global middleware with the auth service
func Init(authService services.AuthServiceInterface) {
	globalMiddleware = NewAuthMiddleware(authService)
}
