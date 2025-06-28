package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrorWrap(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err1 := New("some-error-1")
		err2 := Wrap(err1, "some-error-2")

		assert.NotNil(t, err2.cause)
		assert.True(t, errors.Is(err1, err2.cause))
	})
}

func Test_ErrorUnWrap(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err1 := New("some-error-1")
		err2 := Wrap(err1, "some-error-2")

		assert.Nil(t, err1.Unwrap())
		assert.True(t, errors.Is(err1, err2.Unwrap()))
	})
}

func Test_ErrorIs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err1 := New("some-error-1")
		err2 := Wrap(err1, "some-error-2")

		assert.True(t, Is(err2, err1))
	})
}

func Test_ErrorAs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err1 := New("some-error-1")
		err2 := Wrap(err1, "some-error-2")

		var errT *Error
		assert.True(t, As(err2, &errT))
	})
}
