package apierr

import (
	"github.com/dynamicgo/xerrors"
)

// APIErr .
type APIErr interface {
	error
	Code() int
}

type apiErr struct {
	message string
	code    int
}

// New .
func New(code int, message string) APIErr {
	return &apiErr{
		message: message,
		code:    code,
	}
}

func (err *apiErr) Error() string {
	return err.message
}

func (err *apiErr) Code() int {
	return err.code
}

// ErrInternal .
var ErrInternal = New(-1, "internal error")

// ErrSucceed .
var ErrSucceed = New(0, "error success")

// As convert any err to APIErr
func As(err error) APIErr {

	if err == nil {
		return ErrSucceed
	}

	var apiErr APIErr

	if xerrors.As(err, &apiErr) {
		return apiErr
	}

	return ErrInternal
}
