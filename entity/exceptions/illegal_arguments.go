package exceptions

import (
	"fmt"
	"net/http"
)

type errIllegalArguments struct {
	Status int
	Err    error
}

func IllegalArguments(err error) error {
	return &errIllegalArguments{
		Status: http.StatusBadRequest,
		Err:    err,
	}
}

func (e *errIllegalArguments) Error() string {
	return fmt.Sprintf("illegal arguments: %+v", e.Err)
}

func (e *errIllegalArguments) StatusCode() int {
	return e.Status
}
