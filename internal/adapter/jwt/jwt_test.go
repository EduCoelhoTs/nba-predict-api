package _jwt_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	_jwt "github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/jwt"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	jwtService := _jwt.NewJWTService(privateKey, 1)

	token, err := jwtService.Generate("12345")
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	jwtService := _jwt.NewJWTService(privateKey, 1)

	token, err := jwtService.Generate("12345")
	require.NoError(t, err)
	require.NotEmpty(t, token)

	userID, err := jwtService.Validate(token)
	require.NoError(t, err)
	require.Equal(t, "12345", userID)
}
