package application

import (
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

//go:generate counterfeiter . Restarter

type Restarter interface {
	commandregistry.Command
	ApplicationRestart(app models.Application, orgName string, spaceName string) error
}

type Restart struct {
	ui      terminal.UI
	config  coreconfig.Reader
	starter Starter
	stopper Stopper
	appReq  requirements.ApplicationRequirement
}

func init() {
	commandregistry.Register(&Restart{})
}

func (cmd *Restart) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "restart",
		ShortName:   "rs",
		Description: T("Restart an app"),
		Usage: []string{
			T("CF_NAME restart APP_NAME"),
		},
	}
}

func (cmd *Restart) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires an argument\n\n") + commandregistry.Commands.CommandUsage("restart"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 1)
	}

	cmd.appReq = requirementsFactory.NewApplicationRequirement(fc.Args()[0])

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		requirementsFactory.NewTargetedSpaceRequirement(),
		cmd.appReq,
	}

	return reqs, nil
}

func (cmd *Restart) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config

	//get start for dependency
	starter := commandregistry.Commands.FindCommand("start")
	starter = starter.SetDependency(deps, false)
	cmd.starter = starter.(Starter)

	//get stop for dependency
	stopper := commandregistry.Commands.FindCommand("stop")
	stopper = stopper.SetDependency(deps, false)
	cmd.stopper = stopper.(Stopper)

	return cmd
}

func (cmd *Restart) Execute(c flags.FlagContext) error {
	app := cmd.appReq.GetApplication()
	return cmd.ApplicationRestart(app, cmd.config.OrganizationFields().Name, cmd.config.SpaceFields().Name)
}

func (cmd *Restart) ApplicationRestart(app models.Application, orgName, spaceName string) error {
	stoppedApp, err := cmd.stopper.ApplicationStop(app, orgName, spaceName)
	if err != nil {
		return err
	}

	cmd.ui.Say("")

	_, err = cmd.starter.ApplicationStart(stoppedApp, orgName, spaceName)
	if err != nil {
		return err
	}
	return nil
}
