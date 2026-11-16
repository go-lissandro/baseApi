package middlewre

import (
	"goRest/configs"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// secretJWT é lido sob demanda (uma vez) para não exigir JWT_SECRET já na
// inicialização do pacote, o que dificultava testes.
var secretJWT = sync.OnceValue(func() []byte { return []byte(configs.EnvJWTSecret()) })

// Chaves usadas em c.Locals para expor os dados do token aos handlers
const (
	RoleCtxKey   = "userRole"
	UserIDCtxKey = "userID"
)

func ValidateToken(tokenString string) (Role, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretJWT(), nil
	})

	if err != nil {
		return "", "", err
	}
	if !token.Valid {
		return "", "", jwt.ErrTokenSignatureInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", jwt.ErrInvalidKey
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", "", jwt.ErrInvalidKey
	}

	userID, _ := claims["sub"].(string)

	return Role(role), userID, nil
}

func RoleValidationMiddleware(requiredRoles ...Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "Missing Authorization Header"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		role, userID, err := ValidateToken(tokenString)
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
		c.Locals(RoleCtxKey, role)
		c.Locals(UserIDCtxKey, userID)
		return c.Next()
	}
}

func GenerateToken(userID string, role Role) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": string(role),
		// exp precisa ser numérico (Unix): time.Time era serializado como texto
		// e fazia o jwt/v5 rejeitar todo token gerado aqui.
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretJWT())
}
