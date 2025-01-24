package controller

type ValidationError struct {
	Message   string
	Parameter string
}

type NotFoundError struct {
	Message string
}

func NewValidationError(message string, parameter string) *ValidationError {
	return &ValidationError{
		Message:   message,
		Parameter: parameter,
	}
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}
