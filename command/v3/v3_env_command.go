package v3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccversion"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	"github.com/liamawhite/cli-with-i18n/command/v3/shared"
)

//go:generate counterfeiter . V3EnvActor

type V3EnvActor interface {
	CloudControllerAPIVersion() string
	GetEnvironmentVariablesByApplicationNameAndSpace(appName string, spaceGUID string) (v3action.EnvironmentVariableGroups, v3action.Warnings, error)
}

type V3EnvCommand struct {
	RequiredArgs    flag.EnvironmentArgs `positional-args:"yes"`
	usage           interface{}          `usage:"CF_NAME v3-env APP_NAME"`
	relatedCommands interface{}          `related_commands:"v3-app, v3-apps, v3-set-env, v3-unset-env, running-environment-variable-group, staging-environment-variable-group"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       V3EnvActor
}

func (cmd *V3EnvCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	ccClient, _, err := shared.NewClients(config, ui, true)
	if err != nil {
		if v3Err, ok := err.(ccerror.V3UnexpectedResponseError); ok && v3Err.ResponseCode == http.StatusNotFound {
			return translatableerror.MinimumAPIVersionNotMetError{MinimumVersion: ccversion.MinVersionV3}
		}

		return err
	}
	cmd.Actor = v3action.NewActor(nil, ccClient, config)

	return nil
}

func (cmd V3EnvCommand) Execute(args []string) error {
	cmd.UI.DisplayText(command.ExperimentalWarning)
	cmd.UI.DisplayNewline()

	err := command.MinimumAPIVersionCheck(cmd.Actor.CloudControllerAPIVersion(), ccversion.MinVersionV3)
	if err != nil {
		return err
	}

	err = cmd.SharedActor.CheckTarget(cmd.Config, true, true)
	if err != nil {
		return shared.HandleError(err)
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return shared.HandleError(err)
	}

	appName := cmd.RequiredArgs.AppName
	cmd.UI.DisplayTextWithFlavor("Getting env variables for app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":   appName,
		"OrgName":   cmd.Config.TargetedOrganization().Name,
		"SpaceName": cmd.Config.TargetedSpace().Name,
		"Username":  user.Name,
	})
	cmd.UI.DisplayNewline()

	envGroups, warnings, err := cmd.Actor.GetEnvironmentVariablesByApplicationNameAndSpace(
		appName,
		cmd.Config.TargetedSpace().GUID,
	)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	if len(envGroups.SystemProvided) > 0 || len(envGroups.ApplicationProvided) > 0 {
		cmd.UI.DisplayHeader("System-Provided:")
		cmd.displayEnvGroup(envGroups.SystemProvided)
		if len(envGroups.ApplicationProvided) > 0 {
			cmd.UI.DisplayNewline()
			cmd.displayEnvGroup(envGroups.ApplicationProvided)
		}
	} else {
		cmd.UI.DisplayText("No system-provided env variables have been set")
	}
	cmd.UI.DisplayNewline()

	if len(envGroups.UserProvided) > 0 {
		cmd.UI.DisplayHeader("User-Provided:")
		cmd.displayEnvGroup(envGroups.UserProvided)
	} else {
		cmd.UI.DisplayText("No user-provided env variables have been set")
	}
	cmd.UI.DisplayNewline()

	if len(envGroups.RunningGroup) > 0 {
		cmd.UI.DisplayHeader("Running Environment Variable Groups:")
		cmd.displayEnvGroup(envGroups.RunningGroup)
	} else {
		cmd.UI.DisplayText("No running env variables have been set")
	}
	cmd.UI.DisplayNewline()

	if len(envGroups.StagingGroup) > 0 {
		cmd.UI.DisplayHeader("Staging Environment Variable Groups:")
		cmd.displayEnvGroup(envGroups.StagingGroup)
	} else {
		cmd.UI.DisplayText("No staging env variables have been set")
	}
	cmd.UI.DisplayNewline()

	return nil
}

func (cmd V3EnvCommand) displayEnvGroup(group map[string]interface{}) error {
	for key, val := range group {
		valJSON, err := json.MarshalIndent(val, "", " ")
		if err != nil {
			return err
		}
		cmd.UI.DisplayText(fmt.Sprintf("%s: %s", key, string(valJSON)))
	}

	return nil
}
