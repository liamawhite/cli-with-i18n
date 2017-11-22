package v3

import (
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

//go:generate counterfeiter . V3SetHealthCheckActor

type V3SetHealthCheckActor interface {
	CloudControllerAPIVersion() string
	SetApplicationProcessHealthCheckTypeByNameAndSpace(appName string, spaceGUID string, healthCheckType string, httpEndpoint string, processType string) (v3action.Application, v3action.Warnings, error)
}

type V3SetHealthCheckCommand struct {
	RequiredArgs flag.SetHealthCheckArgs `positional-args:"yes"`
	HTTPEndpoint string                  `long:"endpoint" default:"/" description:"Path on the app"`
	ProcessType  string                  `long:"process" default:"web" description:"App process to update"`
	usage        interface{}             `usage:"CF_NAME v3-set-health-check APP_NAME (process | port | http [--endpoint PATH]) [--process PROCESS]\n\nEXAMPLES:\n   cf v3-set-health-check worker-app process --process worker\n   cf v3-set-health-check my-web-app http --endpoint /foo"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       V3SetHealthCheckActor
}

func (cmd *V3SetHealthCheckCommand) Setup(config command.Config, ui command.UI) error {
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

func (cmd V3SetHealthCheckCommand) Execute(args []string) error {
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

	cmd.UI.DisplayTextWithFlavor("Updating health check type for app {{.AppName}} process {{.ProcessType}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":     cmd.RequiredArgs.AppName,
		"ProcessType": cmd.ProcessType,
		"OrgName":     cmd.Config.TargetedOrganization().Name,
		"SpaceName":   cmd.Config.TargetedSpace().Name,
		"Username":    user.Name,
	})
	cmd.UI.DisplayNewline()

	app, warnings, err := cmd.Actor.SetApplicationProcessHealthCheckTypeByNameAndSpace(
		cmd.RequiredArgs.AppName,
		cmd.Config.TargetedSpace().GUID,
		cmd.RequiredArgs.HealthCheck.Type,
		cmd.HTTPEndpoint,
		cmd.ProcessType,
	)

	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	cmd.UI.DisplayOK()

	if app.Started() {
		cmd.UI.DisplayNewline()
		cmd.UI.DisplayText("TIP: An app restart is required for the change to take effect.")
	}

	return nil
}
