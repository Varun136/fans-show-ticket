package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func (h *authHandler) GenerateJWT(email string) (string, error) {

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
