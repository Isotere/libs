package errors

import stdErrors "errors"

func Is(err, target error) bool {
	return stdErrors.Is(err, target)
}

func As(err error, target any) bool {
	return stdErrors.As(err, &target)
}

func Cause(err error) error {
	for err != nil {
		cause, ok := err.(interface{ Cause() error })
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return err
}
