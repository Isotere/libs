package errors

import "github.com/Isotere/libs/stack"

func (e *Error) WithStacktrace() *Error {
	if e == nil {
		return nil
	}

	e.stacktrace = stack.Callers()

	return e
}
