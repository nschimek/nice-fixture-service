package rest_error

import "net/http"

type ErrorType string

const (
	Authentication ErrorType = "Authentication"
	Authorization ErrorType = "Authorization"
	BadRequest ErrorType = "Bad Request"
	Internal ErrorType = "Internal"
	NotFound ErrorType = "Not Found"
)

var errorMap = map[ErrorType]*Error {
	Authentication: {Type: Authentication, Code: http.StatusUnauthorized, Message: "Unauthenticated - not logged in"},
	Authorization: {Type: Authorization, Code: http.StatusForbidden, Message: "Unauthorized - insuffucient permissions for this action"},
	BadRequest: {Type: BadRequest, Code: http.StatusBadRequest, Message: "Bad request - input field validation failure"},
	Internal: {Type: Internal, Code: http.StatusInternalServerError, Message: "Unexpected internal server error"},
	NotFound: {Type: NotFound, Code: http.StatusNotFound, Message: "Requested content not found"},
}

type Error struct {
	Type ErrorType `json:"type"`
	Code int `json:"code"`
	Message string `json:"message"`
	Details []string `json:"details"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(t ErrorType) *Error {
	return errorMap[t]
}

// Returns a new Internal Server Error (500) with the given error if it's not nil.
func NewInternal(err error) *Error {
	if err != nil {
		e := errorMap[Internal]
		e.Details = []string{err.Error()}
		return e
	}
	return nil
}

func NewWithDetail(t ErrorType, detail string) *Error {
	e := errorMap[t]
	e.Details = []string{detail}
	return e
}

func NewWithDetails(t ErrorType, details []string) *Error {
	e := errorMap[t]
	e.Details = details
	return e
}

