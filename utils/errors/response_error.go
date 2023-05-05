package errors

import "net/http"

type ResErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"errors"`
}

func ResponseError(message string) (err *ResErr) {
	return &ResErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad request",
	}
}

func ResponseNotFound(message string) (err *ResErr) {
	return &ResErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not found",
	}
}

func ResponseServerError(message string) (err *ResErr) {
	return &ResErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
}
