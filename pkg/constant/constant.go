package constant

import "errors"

const (
	// Error messages
	ErrMessageInvalidRequestBody  = "Invalid request body"
	ErrMessageUnauthorized        = "Unauthorized"
	ErrMessageForbidden           = "Forbidden"
	ErrMessageNotFound            = "Resource not found"
	ErrMessageInternalServerError = "Internal server error"
	ErrMessageBadRequest          = "Bad request"
	ErrMessageUserNotFound        = "User not found"
	ErrMessageUserAlreadyExists   = "User with this email already exists"
	ErrMessageInvalidCredentials  = "Invalid credentials"
	ErrMissingCredential          = "missing credential"
	ErrInvalidExpiredToken        = "invalid or expired token"
)

const (
	// Success code
	SuccessCodeOK      = "00"
	SuccessCodeCreated = "01"
	// Error code
	ErrCodeUnauthorized       = "40"
	ErrCodeForbidden          = "41"
	ErrCodeNotFound           = "42"
	ErrCodeBadRequest         = "43"
	ErrCodeInternal           = "44"
	ErrCodeUserNotFound       = "45"
	ErrCodeUserAlreadyExists  = "46"
	ErrCodeInvalidCredentials = "47"
)

const (
	FormFieldRequired   = "this field is required"
	FormFieldInvalid    = "this field is invalid"
	FormFieldErrMinimum = "this field must be at least %s characters long"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrBadRequest         = errors.New("bad request")
	ErrInternal           = errors.New("internal server error")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user with this email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
