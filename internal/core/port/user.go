package port

import "github.com/EduCoelhoTs/nba-predict-api/internal/core/domain"

type UserServiceInterface interface {
	CreateUser(firstName, lastName, email, birthDate, password string) (domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(id, firstName, lastName, email, birthDate, password string) (domain.User, error)
	DeleteUser(id string) error
}

type UserRepositoryInterface interface {
	CreateUser(user domain.User) error
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(id string) error
}
