package v3

import (
	"net/http"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccversion"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	sharedV2 "github.com/liamawhite/cli-with-i18n/command/v2/shared"
	"github.com/liamawhite/cli-with-i18n/command/v3/shared"
)

//go:generate counterfeiter . SetOrgDefaultIsolationSegmentActor

type SetOrgDefaultIsolationSegmentActor interface {
	CloudControllerAPIVersion() string
	GetIsolationSegmentByName(isoSegName string) (v3action.IsolationSegment, v3action.Warnings, error)
	SetOrganizationDefaultIsolationSegment(orgGUID string, isoSegGUID string) (v3action.Warnings, error)
}

//go:generate counterfeiter . SetOrgDefaultIsolationSegmentActorV2

type SetOrgDefaultIsolationSegmentActorV2 interface {
	GetOrganizationByName(orgName string) (v2action.Organization, v2action.Warnings, error)
}

type SetOrgDefaultIsolationSegmentCommand struct {
	RequiredArgs    flag.OrgIsolationArgs `positional-args:"yes"`
	usage           interface{}           `usage:"CF_NAME set-org-default-isolation-segment ORG_NAME SEGMENT_NAME"`
	relatedCommands interface{}           `related_commands:"org, set-space-isolation-segment"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       SetOrgDefaultIsolationSegmentActor
	ActorV2     SetOrgDefaultIsolationSegmentActorV2
}

func (cmd *SetOrgDefaultIsolationSegmentCommand) Setup(config command.Config, ui command.UI) error {
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

	ccClientV2, uaaClientV2, err := sharedV2.NewClients(config, ui, true)
	if err != nil {
		return err
	}
	cmd.ActorV2 = v2action.NewActor(ccClientV2, uaaClientV2, config)

	return nil
}

func (cmd SetOrgDefaultIsolationSegmentCommand) Execute(args []string) error {
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

	cmd.UI.DisplayTextWithFlavor("Setting isolation segment {{.IsolationSegmentName}} to default on org {{.OrgName}} as {{.CurrentUser}}...", map[string]interface{}{
		"IsolationSegmentName": cmd.RequiredArgs.IsolationSegmentName,
		"OrgName":              cmd.RequiredArgs.OrganizationName,
		"CurrentUser":          user.Name,
	})

	org, v2Warnings, err := cmd.ActorV2.GetOrganizationByName(cmd.RequiredArgs.OrganizationName)
	cmd.UI.DisplayWarnings(v2Warnings)
	if err != nil {
		return sharedV2.HandleError(err)
	}

	isoSeg, v3Warnings, err := cmd.Actor.GetIsolationSegmentByName(cmd.RequiredArgs.IsolationSegmentName)
	cmd.UI.DisplayWarnings(v3Warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	v3Warnings, err = cmd.Actor.SetOrganizationDefaultIsolationSegment(org.GUID, isoSeg.GUID)
	cmd.UI.DisplayWarnings(v3Warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	cmd.UI.DisplayOK()
	cmd.UI.DisplayNewline()
	cmd.UI.DisplayText("In order to move running applications to this isolation segment, they must be restarted.")

	return nil
}
