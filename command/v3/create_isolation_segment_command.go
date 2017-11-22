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

//go:generate counterfeiter . CreateIsolationSegmentActor

type CreateIsolationSegmentActor interface {
	CloudControllerAPIVersion() string
	CreateIsolationSegmentByName(isolationSegment v3action.IsolationSegment) (v3action.Warnings, error)
}

type CreateIsolationSegmentCommand struct {
	RequiredArgs    flag.IsolationSegmentName `positional-args:"yes"`
	usage           interface{}               `usage:"CF_NAME create-isolation-segment SEGMENT_NAME\n\nNOTES:\n   The isolation segment name must match the placement tag applied to the Diego cell."`
	relatedCommands interface{}               `related_commands:"enable-org-isolation, isolation-segments"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       CreateIsolationSegmentActor
}

func (cmd *CreateIsolationSegmentCommand) Setup(config command.Config, ui command.UI) error {
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

func (cmd CreateIsolationSegmentCommand) Execute(args []string) error {
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

	cmd.UI.DisplayTextWithFlavor("Creating isolation segment {{.SegmentName}} as {{.CurrentUser}}...", map[string]interface{}{
		"SegmentName": cmd.RequiredArgs.IsolationSegmentName,
		"CurrentUser": user.Name,
	})

	warnings, err := cmd.Actor.CreateIsolationSegmentByName(v3action.IsolationSegment{
		Name: cmd.RequiredArgs.IsolationSegmentName,
	})
	cmd.UI.DisplayWarnings(warnings)
	if _, ok := err.(v3action.IsolationSegmentAlreadyExistsError); ok {
		cmd.UI.DisplayWarning("Isolation segment {{.IsolationSegmentName}} already exists.", map[string]interface{}{
			"IsolationSegmentName": cmd.RequiredArgs.IsolationSegmentName,
		})
	} else if err != nil {
		return err
	}

	cmd.UI.DisplayOK()

	return nil
}
