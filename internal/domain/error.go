package domain

import (
	"fmt"
)

type ErrorCode string

const (
	ErrCodeValidation   ErrorCode = "VALIDATION_ERROR"
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeInternal     ErrorCode = "INTERNAL_ERROR"
	ErrCodeBadRequest   ErrorCode = "BAD_REQUEST"
)

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details string    `json:"details,omitempty"`
}

func (e *Error) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s (%s)", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewError(code ErrorCode, message string, details ...string) *Error {
	err := &Error{Code: code, Message: message}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

func ValidationError(message, details string) *Error {
	return NewError(ErrCodeValidation, message, details)
}

func NotFoundError(message, details string) *Error {
	return NewError(ErrCodeNotFound, message, details)
}

func UnauthorizedError(message, details string) *Error {
	return NewError(ErrCodeUnauthorized, message, details)
}

func InternalError(message, details string) *Error {
	return NewError(ErrCodeInternal, message, details)
}

func BadRequestError(message, details string) *Error {
	return NewError(ErrCodeBadRequest, message, details)
}
