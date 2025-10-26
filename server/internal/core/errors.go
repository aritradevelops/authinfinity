package core

import "net/http"

type ErrorInfo struct {
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
}

// implements error
type HttpError struct {
	Message    string
	StatusCode int
	Info       ErrorInfo
}

var errorMap = map[int]string{
	http.StatusInternalServerError: "Something went wrong!",
	http.StatusNotFound:            "Nothing right here.",
	http.StatusBadRequest:          "The request is invalid",
	http.StatusForbidden:           "You are not authorized to access this resource",
	http.StatusUnprocessableEntity: "The given data does not satisfy the constraints",
}

func (e HttpError) Error() string {
	return e.Message
}

func NewInternalServerError() HttpError {
	return HttpError{
		Message:    errorMap[http.StatusInternalServerError],
		StatusCode: http.StatusInternalServerError,
	}
}
func NewNotFoundError() HttpError {
	return HttpError{
		Message:    errorMap[http.StatusNotFound],
		StatusCode: http.StatusNotFound,
	}
}
func NewBadRequestError() HttpError {
	return HttpError{
		Message:    errorMap[http.StatusBadRequest],
		StatusCode: http.StatusBadRequest,
	}
}

func NewRequestValidationError(errors []ValidationError) HttpError {
	return HttpError{
		Message:    errorMap[http.StatusUnprocessableEntity],
		StatusCode: http.StatusUnprocessableEntity,
		Info: ErrorInfo{
			ValidationErrors: errors,
		},
	}
}
