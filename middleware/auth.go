package middlewre

import (
	"context"
	"goRest/configs"
	"net/http"
	"strings"

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

func RoleValidationMiddleware(requiredRoles ...Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer")

			role, err := ValidateToken(tokenString)

			if err != nil {
				http.Error(w, "Invalid Token", http.StatusUnauthorized)
				return
			}

			authorized := false
			for _, reqRole := range requiredRoles {
				if role == reqRole {
					authorized = true
					break
				}
			}
			if !authorized {
				http.Error(w, "Acess Denied", http.StatusForbidden)
			}
			ctx := context.WithValue(r.Context(), roleCtxKey, role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
