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

func TestApiError(t *testing.T) {
	err := testError(errTest)

	require.True(t, xerrors.Is(err, errTest))

	require.Equal(t, As(err), ErrInternal)

	err = testError(errTest2)

	require.False(t, xerrors.Is(err, errTest))

	require.Equal(t, As(err), errTest2)

	require.Equal(t, As(nil), ErrSucceed)

	err = testError(err)

	println(err.Error())

}
