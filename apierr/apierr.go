package apierr

import (
	"github.com/dynamicgo/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// As convert any err to APIErr
func As(err error, deferr APIErr) APIErr {

	if err == nil {
		panic("invalid input")
	}

	var ae APIErr

	if xerrors.As(err, &ae) {
		return ae
	}

	s, ok := status.FromError(err)

	if ok && s.Code() > 5000 {
		return New(-int(s.Code()-10000), s.Message())
	}

	return deferr
}

// AsGrpcError .
func AsGrpcError(err APIErr) error {

	code := uint32(-err.Code())

	return status.New(codes.Code(10000+code), err.Error()).Err()
}
