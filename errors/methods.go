package errors

import "github.com/Isotere/libs/stack"

func (e *Error) Stacktrace() stack.Trace {
	if e == nil || e.stacktrace == nil {
		return nil
	}

	return e.stacktrace.StackTrace()
}
