package service

import (
	"context"

	"github.com/EduCoelhoTs/nba-predict-api/internal/core/domain"
	"github.com/EduCoelhoTs/nba-predict-api/internal/core/port"
	"github.com/EduCoelhoTs/nba-predict-api/pkg/xuuid"
)

type User struct {
	Repository port.UserRepositoryInterface
}

func NewUserService(ctx context.Context, repository port.UserRepositoryInterface) *User {
	return &User{
		Repository: repository,
	}
}

func (s *User) CreateUser(ctx context.Context, firstName, lastName, email, birthDate, password string) (domain.User, error) {
	id := xuuid.NewV7()
	user := domain.NewUser(
		id,
		firstName,
		lastName,
		email,
		birthDate,
		password,
	)

	if err := user.IsValid(); err != nil {
		return nil, err
	}

	if err := s.Repository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.Repository.GetAllUsers(ctx)
}

func (s *User) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	return s.Repository.GetUserByID(ctx, id)
}

func (s *User) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.Repository.GetUserByEmail(ctx, email)
}

func (s *User) UpdateUser(ctx context.Context, id, firstName, lastName, email, birthDate, password string) (domain.User, error) {
	user := domain.NewUser(
		id,
		firstName,
		lastName,
		email,
		birthDate,
		password,
	)

	if err := user.IsValid(); err != nil {
		return nil, err
	}

	if err := s.Repository.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) DeleteUser(ctx context.Context, id string) error {
	return s.Repository.DeleteUser(ctx, id)
}
