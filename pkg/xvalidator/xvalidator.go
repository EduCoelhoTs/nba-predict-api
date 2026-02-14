package xvalidator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateStruct(s any) error {
	var validationErrors validator.ValidationErrors
	if err := validate.Struct(s); err != nil {
		if errors.As(err, &validationErrors) {
			var ResultError error
			for _, err := range validationErrors {
				errMsg := fmt.Errorf("field %s is invalid. Field must be: %s", err.Field(), err.ActualTag())
				ResultError = errors.Join(ResultError, errMsg)
			}
			return ResultError
		}
		return err
	}

	return nil
}
