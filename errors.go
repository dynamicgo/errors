package xerrors

import "reflect"

// Error .
type Error interface {
	error              // mixin standard error interface
	CallStack() string // get call stack
	Raised() error     // error chain
}

// PrintStack print stack flag
var PrintStack = true

type facadeImpl struct {
}

func (facade *facadeImpl) Is(err, target error) bool {

	current := target

	for {
		if current == err {
			return true
		}

		e, ok := err.(Error)

		if !ok {
			return false
		}

		current = e.Raised()

		if current == nil {
			return false
		}
	}

}

func (facade *facadeImpl) As(err interface{}, target error) bool {

	errPointT := reflect.TypeOf(err)

	if errPointT.Kind() != reflect.Ptr {
		panic("invalid type")
	}

	errT := errPointT.Elem()

	if errT.Kind() != reflect.Ptr || errT.Kind() != reflect.Interface {
		panic("invalid type")
	}

	current := target

	for {
		currentT := reflect.TypeOf(current)

		if currentT == errT || currentT.Implements(errT) {
			reflect.ValueOf(err).Elem().Set(reflect.ValueOf(current))
			return true
		}

		e, ok := err.(Error)

		if !ok {
			return false
		}

		current = e.Raised()

		if current == nil {
			return false
		}
	}
}

func init() {
	RegisterFacade("xerrors", &facadeImpl{})
}
