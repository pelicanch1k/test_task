package validation

import "github.com/go-playground/validator/v10"

var validatorInstance *validator.Validate

func GetValidator() *validator.Validate {
	if validatorInstance == nil {
		validatorInstance = validator.New()
	}
	return validatorInstance
}