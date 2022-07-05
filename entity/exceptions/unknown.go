package exceptions

import (
	"fmt"
	"net/http"
)

type errUnknown struct {
	Status int
	Err    error
}

func (e *errUnknown) Error() string {
	return fmt.Sprintf("unknown error: %+v", e.Err)
}

func Unknown(err error) error {
	return &errUnknown{
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}

func (e *errUnknown) StatusCode() int {
	return e.Status
}
