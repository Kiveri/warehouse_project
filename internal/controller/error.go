package controller

type ValidationError struct {
	Message   string
	Parameter string
}

func NewValidationError(message string, parameter string) *ValidationError {
	return &ValidationError{
		Message:   message,
		Parameter: parameter,
	}
}
