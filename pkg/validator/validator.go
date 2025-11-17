package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

// Providers is a Wire provider set that provides a new validator.
var Providers = wire.NewSet(NewValidator)

// NewValidator creates a new validator instance.
func NewValidator() *validator.Validate {
	return validator.New()
}
