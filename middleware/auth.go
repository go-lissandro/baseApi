package middlewre

import (
	"goRest/configs"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

var SecretJWT = []byte(configs.EnvJWTSecret())

type ctxKey string

const roleCtxKey ctxKey = "userRole"

func ValidateToken(tokenString string) (Role, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return SecretJWT, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	return Role(role), nil
}

func RoleValidationMiddleware(requiredRoles ...Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "Missing Authorization Header"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		role, err := ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "Invalid Token"})
		}

		authorized := false
		for _, reqRole := range requiredRoles {
			if role == reqRole {
				authorized = true
				break
			}
		}

		if !authorized {
			return c.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"error": "Access Denied"})
		}
		c.Locals(roleCtxKey, role)
		return c.Next()
	}
}
