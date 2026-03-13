package _jwt

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret  *ecdsa.PrivateKey
	hourExp int
}

type Claims struct {
	UserID string `json:"sub"`
	jwt.RegisteredClaims
}

func NewJWTService(secret *ecdsa.PrivateKey, hoursExp int) *JWTService {
	return &JWTService{secret: secret, hourExp: hoursExp}
}

func (s *JWTService) Generate(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.hourExp) * time.Hour)),
		},
	})

	return token.SignedString(s.secret)
}

func (s *JWTService) Validate(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			return &s.secret.PublicKey, nil
		},
	)

	if err != nil {
		return "", fmt.Errorf("adapter.jwt.validate error: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", fmt.Errorf("adapter.jwt.validate error: invalid token")

}
