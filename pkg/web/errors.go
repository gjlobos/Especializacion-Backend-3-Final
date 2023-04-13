package web

import "net/http"

type ErrorApi struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorApi) Error() string {
	return e.Message
}

func NewNotFoundApiError(message string) error {
	return &ErrorApi{http.StatusNotFound, "not_found", message}
}

func NewBadRequestApiError(message string) error {
	return &ErrorApi{http.StatusBadRequest, "bad_request", message}
}

func NewInternalServerApiError(message string) error {
	return &ErrorApi{http.StatusInternalServerError, "internal_server_error", message}
}
