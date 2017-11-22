package v2

import (
	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/v2/shared"
	"github.com/liamawhite/cli-with-i18n/util/ui"
)

//go:generate counterfeiter . SpacesActor

type SpacesActor interface {
	GetOrganizationSpaces(orgGUID string) ([]v2action.Space, v2action.Warnings, error)
}

type SpacesCommand struct {
	usage           interface{} `usage:"CF_NAME spaces"`
	relatedCommands interface{} `related_commands:"target"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       SpacesActor
}

func (cmd *SpacesCommand) Setup(config command.Config, ui command.UI) error {
	cmd.Config = config
	cmd.UI = ui
	cmd.SharedActor = sharedaction.NewActor(config)

	ccClient, _, err := shared.NewClients(config, ui, true)
	if err != nil {
		return err
	}
	cmd.Actor = v2action.NewActor(ccClient, nil, config)

	return nil
}

func (cmd SpacesCommand) Execute([]string) error {
	err := cmd.SharedActor.CheckTarget(cmd.Config, true, false)
	if err != nil {
		return shared.HandleError(err)
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return shared.HandleError(err)
	}

	cmd.UI.DisplayTextWithFlavor("Getting spaces in org {{.OrgName}} as {{.CurrentUser}}...", map[string]interface{}{
		"OrgName":     cmd.Config.TargetedOrganization().Name,
		"CurrentUser": user.Name,
	})
	cmd.UI.DisplayNewline()

	spaces, warnings, err := cmd.Actor.GetOrganizationSpaces(cmd.Config.TargetedOrganization().GUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	if len(spaces) == 0 {
		cmd.UI.DisplayText("No spaces found.")
	} else {
		cmd.displaySpaces(spaces)
	}

	return nil
}

func (cmd SpacesCommand) displaySpaces(spaces []v2action.Space) {
	table := [][]string{{cmd.UI.TranslateText("name")}}
	for _, space := range spaces {
		table = append(table, []string{space.Name})
	}
	cmd.UI.DisplayTableWithHeader("", table, ui.DefaultTableSpacePadding)
}
