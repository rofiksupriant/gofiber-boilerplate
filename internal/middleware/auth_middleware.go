package middleware

import (
	"boilerplate/internal/security"
	"boilerplate/internal/usecase"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func NewAuth(userUserCase *usecase.UserUseCase) fiber.Handler {
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

		// decrypt token
		tokenPlain, err := security.Decrypt(strings.TrimPrefix(bearerToken, prefix))
		if err != nil {
			log.Errorf("error decrypt token %v", err)
			return fiber.ErrUnauthorized
		}

		// verify token
		claims, ok := security.VerifyToken(*tokenPlain)
		if !ok {
			log.Warnf("faild varify token %v", err)
			return fiber.ErrUnauthorized
		}

		auth, err := userUserCase.Verify(claims.Subject)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		c.Locals("auth", auth)

		return c.Next()
	}
}
