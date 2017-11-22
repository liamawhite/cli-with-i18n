package errors

import (
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
)

type ModelNotFoundError struct {
	ModelType string
	ModelName string
}

func NewModelNotFoundError(modelType, name string) error {
	return &ModelNotFoundError{
		ModelType: modelType,
		ModelName: name,
	}
}

func (err *ModelNotFoundError) Error() string {
	return err.ModelType + " " + err.ModelName + T(" not found")
}
