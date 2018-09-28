package apierr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dynamicgo/xerrors"
)

func testError(err error) error {
	return testError2(err)
}

func testError2(err error) error {
	return xerrors.Wrapf(err, "test stack error")
}

var errTest = errors.New("test error")
var errTest2 = New(1, "test")
var errInternal = New(-1, "internal error")

func TestApiError(t *testing.T) {
	err := testError(errTest)

	require.True(t, xerrors.Is(err, errTest))

	require.Equal(t, As(err, errInternal), errInternal)

	err = testError(errTest2)

	require.False(t, xerrors.Is(err, errTest))

	require.Equal(t, As(err, errInternal), errTest2)

	err = testError(err)

	println(err.Error())

}
