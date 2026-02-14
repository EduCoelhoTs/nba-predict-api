package port

import (
	"context"

	"github.com/EduCoelhoTs/nba-predict-api/internal/core/domain"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, firstName, lastName, email, birthDate, password string) (domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, id, firstName, lastName, email, birthDate, password string) (domain.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id string) error
}
