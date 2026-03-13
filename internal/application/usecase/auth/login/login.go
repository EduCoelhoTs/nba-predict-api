package loginusecase

import (
	"context"

	authport "github.com/EduCoelhoTs/base-hex-arq-api/internal/core/port/auth"
	userport "github.com/EduCoelhoTs/base-hex-arq-api/internal/core/port/user"
	"github.com/EduCoelhoTs/base-hex-arq-api/pkg/xcrypto"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginUseCase interface {
	Execute(ctx context.Context, input LoginInput) (string, error)
}

type loginUseCase struct {
	tokenService   authport.TokenService
	userRepository userport.UserRepositoryInterface
}

func NewLoginUseCase(tokenService authport.TokenService, userRepository userport.UserRepositoryInterface) *loginUseCase {
	return &loginUseCase{
		tokenService:   tokenService,
		userRepository: userRepository,
	}
}

func (uc *loginUseCase) Execute(ctx context.Context, input LoginInput) (string, error) {
	user, err := uc.userRepository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if err := xcrypto.ComparePassword(user.GetPassword(), input.Password); err != nil {
		return "", err
	}

	token, err := uc.tokenService.Generate(user.GetID())
	if err != nil {
		return "", err
	}

	return token, nil
}
