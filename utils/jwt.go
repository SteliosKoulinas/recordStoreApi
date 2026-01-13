package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "CHANGE_ME_TO_A_SECURE_KEY"
	}
	return []byte(key)
}

var jwtKey = getJWTKey()

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
