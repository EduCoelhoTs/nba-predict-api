package bootstrap

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	_jwt "github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/jwt"
	"github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/repository/postgres"
	loginusecase "github.com/EduCoelhoTs/base-hex-arq-api/internal/application/usecase/auth/login"
	authcontroller "github.com/EduCoelhoTs/base-hex-arq-api/internal/infra/controller/auth"
)

func NewAuthModule(ctx context.Context, db postgres.QueriesPort) authcontroller.Controller {
	repository := postgres.NewUserRepository(db)
	tokenService := _jwt.NewJWTService(&ecdsa.PrivateKey{
		ecdsa.PublicKey{},
		&big.Int{},
	}, 8)
	usecase := loginusecase.NewLoginUseCase(tokenService, repository)

	return authcontroller.NewController(usecase)
}
