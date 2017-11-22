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

//go:generate counterfeiter . EnableOrgIsolationActor

type EnableOrgIsolationActor interface {
	CloudControllerAPIVersion() string
	EntitleIsolationSegmentToOrganizationByName(isolationSegmentName string, orgName string) (v3action.Warnings, error)
}

type EnableOrgIsolationCommand struct {
	RequiredArgs    flag.OrgIsolationArgs `positional-args:"yes"`
	usage           interface{}           `usage:"CF_NAME enable-org-isolation ORG_NAME SEGMENT_NAME"`
	relatedCommands interface{}           `related_commands:"create-isolation-segment, isolation-segments, set-org-default-isolation-segment, set-space-isolation-segment"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       EnableOrgIsolationActor
}

func (cmd *EnableOrgIsolationCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	client, _, err := shared.NewClients(config, ui, true)
	if err != nil {
		if v3Err, ok := err.(ccerror.V3UnexpectedResponseError); ok && v3Err.ResponseCode == http.StatusNotFound {
			return translatableerror.MinimumAPIVersionNotMetError{MinimumVersion: ccversion.MinVersionIsolationSegmentV3}
		}

		return err
	}
	cmd.Actor = v3action.NewActor(nil, client, config)

	return nil
}

func (cmd EnableOrgIsolationCommand) Execute(args []string) error {
	err := command.MinimumAPIVersionCheck(cmd.Actor.CloudControllerAPIVersion(), ccversion.MinVersionIsolationSegmentV3)
	if err != nil {
		return err
	}

	err = cmd.SharedActor.CheckTarget(cmd.Config, false, false)
	if err != nil {
		return shared.HandleError(err)
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Enabling isolation segment {{.SegmentName}} for org {{.OrgName}} as {{.CurrentUser}}...", map[string]interface{}{
		"SegmentName": cmd.RequiredArgs.IsolationSegmentName,
		"OrgName":     cmd.RequiredArgs.OrganizationName,
		"CurrentUser": user.Name,
	})

	warnings, err := cmd.Actor.EntitleIsolationSegmentToOrganizationByName(cmd.RequiredArgs.IsolationSegmentName, cmd.RequiredArgs.OrganizationName)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	cmd.UI.DisplayOK()

	return nil
}
