package utils

import (
	"time"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, secret string) (string, error) {
	log.Println("[DEBUG] secret:", secret)
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}