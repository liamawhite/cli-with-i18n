package fakecommand

import (
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
)

type FakeCommand3 struct {
}

func init() {
	commandregistry.Register(FakeCommand3{})
}

func (cmd FakeCommand3) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name: "this-is-a-really-long-command-name-123123123123123123123",
	}
}

func (cmd FakeCommand3) Requirements(_ requirements.Factory, _ flags.FlagContext) ([]requirements.Requirement, error) {
	return []requirements.Requirement{}, nil
}

func (cmd FakeCommand3) SetDependency(deps commandregistry.Dependency, _ bool) commandregistry.Command {
	return cmd
}

func (cmd FakeCommand3) Execute(c flags.FlagContext) error {
	return nil
}
