package xuuid

import "github.com/google/uuid"

func NewV7() string {
	result, _ := uuid.NewV7()
	return result.String()
}
