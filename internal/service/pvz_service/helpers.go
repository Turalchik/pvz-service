package pvz_service

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID string
	Role   string
	jwt.RegisteredClaims
}

func generateToken(secretKeyForJWT []byte, userID string, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKeyForJWT)
}
