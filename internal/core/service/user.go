package service

import (
	"github.com/EduCoelhoTs/nba-predict-api/internal/core/domain"
	"github.com/EduCoelhoTs/nba-predict-api/internal/core/port"
	"github.com/EduCoelhoTs/nba-predict-api/pkg/xuuid"
)

type User struct {
	Repository port.UserRepositoryInterface
}

func NewUserService(repository port.UserRepositoryInterface) *User {
	return &User{
		Repository: repository,
	}
}

func (s *User) CreateUser(firstName, lastName, email, birthDate, password string) (domain.User, error) {
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

	if err := s.Repository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) GetAllUsers() ([]domain.User, error) {
	return s.Repository.GetAllUsers()
}

func (s *User) GetUserByID(id string) (domain.User, error) {
	return s.Repository.GetUserByID(id)
}

func (s *User) GetUserByEmail(email string) (domain.User, error) {
	return s.Repository.GetUserByEmail(email)
}

func (s *User) UpdateUser(id, firstName, lastName, email, birthDate, password string) (domain.User, error) {
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

	if err := s.Repository.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) DeleteUser(id string) error {
	return s.Repository.DeleteUser(id)
}
