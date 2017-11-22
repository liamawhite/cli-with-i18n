package applicationfakes

import (
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
)

type FakeAppDisplayer struct {
	AppToDisplay models.Application
	OrgName      string
	SpaceName    string
}

func (displayer *FakeAppDisplayer) ShowApp(app models.Application, orgName, spaceName string) error {
	displayer.AppToDisplay = app
	return nil
}

func (displayer *FakeAppDisplayer) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{Name: "app"}
}

func (displayer *FakeAppDisplayer) SetDependency(_ commandregistry.Dependency, _ bool) commandregistry.Command {
	return displayer
}

func (displayer *FakeAppDisplayer) Requirements(_ requirements.Factory, _ flags.FlagContext) ([]requirements.Requirement, error) {
	return []requirements.Requirement{}, nil
}

func (displayer *FakeAppDisplayer) Execute(_ flags.FlagContext) error {
	return nil
}
