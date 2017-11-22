package errors

import (
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
)

type EmptyDirError struct {
	dir string
}

func NewEmptyDirError(dir string) error {
	return &EmptyDirError{dir: dir}
}

func (err *EmptyDirError) Error() string {
	return err.dir + T(" is empty")
}
