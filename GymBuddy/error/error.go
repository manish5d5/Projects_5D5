package apperror

import "fmt"

// AppError holds a human readable message and an HTTP-like status code.
type AppError struct {
	Message string
	Code    int
	Err     error // underlying error (optional)
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Helper to wrap underlying errors (DB or service errors)
func Wrap(msg string, code int, err error) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
		Err:     err,
	}
}

// Predefined common errors
var (
	ErrNotFound = &AppError{
		Message: "record not found",
		Code:    404,
	}

	ErrInsertFailed = &AppError{
		Message: "failed to insert record",
		Code:    500,
	}

	ErrUpdateFailed = &AppError{
		Message: "failed to update record",
		Code:    500,
	}

	ErrDeleteFailed = &AppError{
		Message: "failed to delete record",
		Code:    500,
	}

	ErrInvalidInput = &AppError{
		Message: "invalid input",
		Code:    400,
	}
)
