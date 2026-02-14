package domain

import "github.com/EduCoelhoTs/nba-predict-api/pkg/xvalidator"

type user struct {
	ID        string `validate:"required,uuid"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	BirthDate string `validate:"required,datetime=2006-01-02"`
	Password  string `validate:"required,min=8"` // Password must be at least 8 characters long
}

type User interface {
	GetID() string
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	GetBirthDate() string
	GetPassword() string
	IsValid() error
}

func NewUser(id, firstName, lastName, email, birthDate, password string) *user {
	return &user{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		BirthDate: birthDate,
		Password:  password,
	}
}

func (u *user) GetID() string {
	return u.ID
}

func (u *user) GetFirstName() string {
	return u.FirstName
}

func (u *user) GetLastName() string {
	return u.LastName
}

func (u *user) GetEmail() string {
	return u.Email
}

func (u *user) GetBirthDate() string {
	return u.BirthDate
}

func (u *user) GetPassword() string {
	return u.Password
}

func (u *user) IsValid() error {
	return xvalidator.ValidateStruct(u)
}
