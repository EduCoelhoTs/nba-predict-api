package xuuid

import "github.com/google/uuid"

func NewV7() string {
	result, _ := uuid.NewV7()
	return result.String()
}

func UUIDFromString(s string) (uuid.UUID, error) {
	result, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, err
	}
	return result, nil
}
