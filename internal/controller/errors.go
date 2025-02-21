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

type NotFoundError struct {
	Message string `json:"message"`
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}
