package controller

import (
	"encoding/json"
	"net/http"
)

// Respond отправляет JSON-ответ клиенту с указанным кодом статуса.
func Respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// ValidationErrorRespond отправляет JSON-ответ с ошибкой.
func ValidationErrorRespond(w http.ResponseWriter, validationError *ValidationError) {
	Respond(w, http.StatusBadRequest, validationError)
}

// InternalServerErrorRespond отправляет JSON-ответ с сообщением об ошибке сервера.
func InternalServerErrorRespond(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func NotFoundErrorRespond(w http.ResponseWriter, notFoundError *NotFoundError) {
	Respond(w, http.StatusNotFound, notFoundError)
}

// EncodeResponse записывает данные из структуры в json и возвращает в качестве ответа
func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := encoder.Encode(response)
	if err != nil {
		ValidationErrorRespond(w, NewValidationError("encode", "error writing json encoding"))

		return err
	}

	return nil
}
