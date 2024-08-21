package urlshortener

import "fmt"

// ErrorType defines the type of error
type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = "VALIDATION_ERROR"

	// ErrorTypeDatabase represents database errors
	ErrorTypeDatabase ErrorType = "DATABASE_ERROR"

	// ErrorTypeNotFound represents not found errors
	ErrorTypeNotFound ErrorType = "NOT_FOUND_ERROR"

	// ErrorTypeInternal represents internal errors
	ErrorTypeInternal ErrorType = "INTERNAL_ERROR"
)

// AppError represents a custom error for the application
type AppError struct {
	Type    ErrorType
	Message string
}

// NewAppError creates a new AppError
func NewAppError(errorType ErrorType, message string) *AppError {
	return &AppError{
		Type:    errorType,
		Message: message,
	}
}

// Error implements the error interface for AppError
func (e *AppError) Error() string {
	return fmt.Sprintf("Type: %s, Message: %s", e.Type, e.Message)
}

// IsOfType checks if the error is of a specific type
func (e *AppError) IsOfType(errorType ErrorType) bool {
	return e.Type == errorType
}
