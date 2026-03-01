package sqlc

import (
	"github.com/EduCoelhoTs/nba-predict-api/internal/core/domain"
	"github.com/EduCoelhoTs/nba-predict-api/pkg/xdate"
)

func (au *AuthUser) ToDomain() domain.User {
	return domain.NewUser(
		au.ID.String(),
		au.FirstName,
		au.LastName,
		au.Email,
		au.BirthDate.Format(xdate.DefaultDateLayout),
		au.Password,
	)
}
