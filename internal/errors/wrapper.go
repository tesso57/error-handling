package errors

import "github.com/cockroachdb/errors"

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func Wrapf(err error, format string, args ...any) error {
	return errors.Wrapf(err, format, args...)
}

func New(message string) error {
	return errors.New(message)
}

func Newf(format string, args ...any) error {
	return errors.Newf(format, args...)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
