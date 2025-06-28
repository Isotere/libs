package errors

import (
	"fmt"

	"github.com/Isotere/libs/stack"
)

// Warning: за счет расширенной структуры хранения (стек и cause) дает оверхед по памяти (x3 - 16B vs 48B) и небольшой оверхед по скорости (+- 3ns на операцию)
// Если нужен только стек, можно рассмотреть вызов Callers в том месте, где нужен стек и обработать этот вызов

type Error struct {
	message    string
	stacktrace *stack.Stack
	cause      error
}

func New(message string) *Error {
	return &Error{
		message: message,
	}
}

func Errorf(format string, args ...interface{}) *Error {
	return New(fmt.Sprintf(format, args...))
}

func Wrap(err error, msg string) *Error {
	return &Error{
		message: msg,
		cause:   err,
	}
}

func Wrapf(err error, format string, args ...interface{}) *Error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.cause != nil {
		return fmt.Sprintf("%s\n %v", e.message, e.cause)
	}

	return fmt.Sprintf("%s", e.message)
}

func (e *Error) Cause() error {
	if e == nil {
		return nil
	}

	return e.cause
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.cause
}
