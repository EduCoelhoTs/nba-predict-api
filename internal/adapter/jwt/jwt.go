package _jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret  string
	hourExp int
}

type Claims struct {
	UserID string `json:"sub"`
	jwt.RegisteredClaims
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: secret}
}

func (s *JWTService) Generate(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.hourExp) * time.Hour)),
		},
	})

	return token.SignedString([]byte(s.secret))
}

func (s *JWTService) Validate(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		Claims{},
		func(t *jwt.Token) (any, error) {
			return []byte(s.secret), nil
		},
	)

	if err != nil {
		return "", fmt.Errorf("adapter.jwt.validate error: %w", err)
	}

	if claims, ok := token.Claims.(Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", fmt.Errorf("adapter.jwt.validate error: invalid token")

}
