package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidatorErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(w http.ResponseWriter, v interface{}) bool {
	err := validate.Struct(v)
	if err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		errs := make([]ValidatorErr, 0)
		for _, fieldError := range validationErrors {
			field := strings.ToLower(fieldError.Field())
			message := getValidationMessage(field, fieldError.Tag(), fieldError.Param())
			errs = append(errs, ValidatorErr{field, message})
		}
		ValidationError(w, "Validation failed!!", errs)
		return false
	}
	return true
}

func getValidationMessage(field string, fieldTag string, param string) string {
	switch fieldTag {
	case "required":
		return field + " is required"
	case "max":
		return field + " must be greater than " + fieldTag
	case "min":
		return field + " must be greater than " + fieldTag
	case "email":
		return field + " must be a valid email"
	case "full_name":
		return field + " must be a valid full name"
	case "mobile":
		return field + " must be a valid mobile"
	case "alphanum":
		return field + " must be a valid username"
	default:
		return "default value is required"

	}
}
