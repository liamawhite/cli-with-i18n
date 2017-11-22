package cfnetworkingaction

import "github.com/liamawhite/cli-with-i18n/actor/v3action"

//go:generate counterfeiter . V3Actor
type V3Actor interface {
	GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	GetApplicationsBySpace(spaceGUID string) ([]v3action.Application, v3action.Warnings, error)
}
