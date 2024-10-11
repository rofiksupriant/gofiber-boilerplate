package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
	"strings"
)

func NewAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearerToken := c.Get(fiber.HeaderAuthorization, "NOT_FOUND")
		if bearerToken == "NOT_FOUND" {
			log.Info("Token not found")
			return fiber.ErrUnauthorized
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(bearerToken, prefix) {
			log.Info("Token not bearer")
			return fiber.ErrUnauthorized
		}

		token := strings.TrimPrefix(bearerToken, prefix)
		log.Info(token)

		return c.Next()
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Error("auth header is null or blank")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized"))
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Errorf("auth header is invalid: %v", authHeader)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
