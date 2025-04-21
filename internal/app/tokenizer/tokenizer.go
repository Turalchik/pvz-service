package tokenizer

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID string
	Role   string
	jwt.RegisteredClaims
}

type TokenizerInterface interface {
	GenerateToken(userID string, role string) (string, error)
	VerifyToken(tokenString string) (*Claims, error)
}

func NewTokenizer(secretKeyForJWT []byte) TokenizerInterface {
	tokenizer := &Tokenizer{
		secretKeyForJWT: make([]byte, len(secretKeyForJWT)),
	}
	copy(tokenizer.secretKeyForJWT, secretKeyForJWT)
	return tokenizer
}

type Tokenizer struct {
	secretKeyForJWT []byte
}

func (tokenizer *Tokenizer) GenerateToken(userID string, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tokenizer.secretKeyForJWT)
}

func (tokenizer *Tokenizer) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenizer.secretKeyForJWT, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
