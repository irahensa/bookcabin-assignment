package customErrors

import (
	"errors"
	"net/http"
)

type CustomError struct {
	Name     string
	Message  string
	Code     string
	Err      error
	HTTPCode int
}

func (e CustomError) Error() string {
	return e.Err.Error()
}

var ErrorDB = CustomError{
	Name:     "Error DB",
	Message:  "Error query data",
	Code:     "5001",
	HTTPCode: http.StatusInternalServerError,
}

var ErrorParsing = CustomError{
	Name:     "error parsing",
	Message:  "Unable to parse data",
	Code:     "5001",
	HTTPCode: http.StatusBadRequest,
}

var DuplicateData = CustomError{
	Name:     "duplicate data",
	Message:  "voucher for this flight already generated",
	Code:     "5001",
	HTTPCode: http.StatusBadRequest,
}

func New(c CustomError, e error) error {
	c.Err = e
	return c
}

func Parse(e error) (c CustomError) {
	c, success := e.(CustomError)
	if success {
		return c
	} else {
		return CustomError{
			Name:     "unknown error",
			Message:  "unknown error",
			HTTPCode: http.StatusInternalServerError,
			Err:      errors.New(e.Error()),
		}
	}
}
