package commandregistry

import (
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
)

//go:generate counterfeiter . Command

type Command interface {
	MetaData() CommandMetadata
	SetDependency(deps Dependency, pluginCall bool) Command
	Requirements(requirementsFactory requirements.Factory, context flags.FlagContext) ([]requirements.Requirement, error)
	Execute(context flags.FlagContext) error
}

type CommandMetadata struct {
	Name            string
	ShortName       string
	Usage           []string
	Description     string
	Flags           map[string]flags.FlagSet
	SkipFlagParsing bool
	TotalArgs       int //Optional: number of required arguments to skip for flag verification
	Hidden          bool
	Examples        []string
}
