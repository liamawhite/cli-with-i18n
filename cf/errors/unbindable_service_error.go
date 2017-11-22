package errors

import (
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
)

type UnbindableServiceError struct {
}

func NewUnbindableServiceError() error {
	return &UnbindableServiceError{}
}

func (err *UnbindableServiceError) Error() string {
	return T("This service doesn't support creation of keys.")
}
