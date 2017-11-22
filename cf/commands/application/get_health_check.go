package application

import (
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf/api/applications"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"

	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type GetHealthCheck struct {
	ui      terminal.UI
	config  coreconfig.Reader
	appReq  requirements.ApplicationRequirement
	appRepo applications.Repository
}

func init() {
	commandregistry.Register(&GetHealthCheck{})
}

func (cmd *GetHealthCheck) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "get-health-check",
		Description: T("Get the health_check_type value of an app"),
		Usage: []string{
			T("CF_NAME get-health-check APP_NAME"),
		},
	}
}

func (cmd *GetHealthCheck) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires APP_NAME as argument\n\n") + commandregistry.Commands.CommandUsage("get-health-check"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 1)
	}

	cmd.ui.Say(T("Getting health_check_type value for ") + terminal.EntityNameColor(fc.Args()[0]))
	cmd.ui.Say("")

	cmd.appReq = requirementsFactory.NewApplicationRequirement(fc.Args()[0])

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		requirementsFactory.NewTargetedSpaceRequirement(),
		cmd.appReq,
	}

	return reqs, nil
}

func (cmd *GetHealthCheck) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.appRepo = deps.RepoLocator.GetApplicationRepository()
	return cmd
}

func (cmd *GetHealthCheck) Execute(fc flags.FlagContext) error {
	app := cmd.appReq.GetApplication()

	cmd.ui.Say(T("health_check_type is ") + terminal.HeaderColor(app.HealthCheckType))
	return nil
}
