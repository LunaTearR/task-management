package pkg

import (
	"fmt"
	"time"
)

type AppError struct {
	Message   string
	Code      int
	Timestamp time.Time
	Err       error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[รหัส: %d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[รหัส: %d] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}
func (e *AppError) WithError(err error) *AppError {
	e.Err = err
	return e
}

func (e *AppError) WithMessage(message string) *AppError {
	e.Message = message
	return e
}

func (e *AppError) WithCode(code int) *AppError {
	e.Code = code
	return e
}
