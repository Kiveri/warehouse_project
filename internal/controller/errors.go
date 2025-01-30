package controller

type ValidationError struct {
	Message   string `json:"message"`
	Parameter string `json:"parameter"`
}

func NewValidationError(message, parameter string) *ValidationError {
	return &ValidationError{
		Message:   message,
		Parameter: parameter,
	}
}
