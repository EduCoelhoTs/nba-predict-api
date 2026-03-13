package authport

type TokenService interface {
	Validate(token string) (string, error)
	Generate(userID string) (string, error)
}
