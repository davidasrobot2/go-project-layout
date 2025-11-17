package response

import (
	"davidasrobot/project-layout/pkg/constant"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ErrorResponse represents a single validation error.
type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// FormatValidationErrors converts validator.ValidationErrors into a slice of ErrorResponse.
func FormatValidationErrors(err error) []ErrorResponse {
	var errors []ErrorResponse

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			errors = append(errors, ErrorResponse{
				Field:   fieldErr.Field(),
				Message: messageForTag(fieldErr),
			})
		}
	}

	return errors
}

// messageForTag returns a user-friendly message for a given validation tag.
func messageForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return constant.FormFieldRequired
	case "min":
		return fmt.Sprintf(constant.FormFieldErrMinimum, fe.Param())
	default:
		return constant.FormFieldInvalid
	}
}
