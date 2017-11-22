package environmentvariablegroup

import (
	"errors"
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf/api/environmentvariablegroups"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	cf_errors "github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type SetRunningEnvironmentVariableGroup struct {
	ui                           terminal.UI
	config                       coreconfig.ReadWriter
	environmentVariableGroupRepo environmentvariablegroups.Repository
}

func init() {
	commandregistry.Register(&SetRunningEnvironmentVariableGroup{})
}

func (cmd *SetRunningEnvironmentVariableGroup) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "set-running-environment-variable-group",
		Description: T("Pass parameters as JSON to create a running environment variable group"),
		ShortName:   "srevg",
		Usage: []string{
			T(`CF_NAME set-running-environment-variable-group '{"name":"value","name":"value"}'`),
		},
	}
}

func (cmd *SetRunningEnvironmentVariableGroup) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires an argument\n\n") + commandregistry.Commands.CommandUsage("set-running-environment-variable-group"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 1)
	}

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
	}
	return reqs, nil
}

func (cmd *SetRunningEnvironmentVariableGroup) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.environmentVariableGroupRepo = deps.RepoLocator.GetEnvironmentVariableGroupsRepository()
	return cmd
}

func (cmd *SetRunningEnvironmentVariableGroup) Execute(c flags.FlagContext) error {
	cmd.ui.Say(T("Setting the contents of the running environment variable group as {{.Username}}...", map[string]interface{}{
		"Username": terminal.EntityNameColor(cmd.config.Username())}))

	err := cmd.environmentVariableGroupRepo.SetRunning(c.Args()[0])
	if err != nil {
		suggestionText := ""

		httpError, ok := err.(cf_errors.HTTPError)
		if ok && httpError.ErrorCode() == cf_errors.MessageParseError {
			suggestionText = T(`

Your JSON string syntax is invalid.  Proper syntax is this:  cf set-running-environment-variable-group '{"name":"value","name":"value"}'`)
		}
		return errors.New(err.Error() + suggestionText)
	}

	cmd.ui.Ok()
	return nil
}
