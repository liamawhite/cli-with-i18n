package v2

import (
	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/v2/shared"
	"github.com/liamawhite/cli-with-i18n/util/ui"
)

//go:generate counterfeiter . OrgsActor

type OrgsActor interface {
	GetOrganizations() ([]v2action.Organization, v2action.Warnings, error)
}

type OrgsCommand struct {
	usage interface{} `usage:"CF_NAME orgs"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       OrgsActor
}

func (cmd *OrgsCommand) Setup(config command.Config, ui command.UI) error {
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

func (cmd OrgsCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(cmd.Config, false, false)
	if err != nil {
		return shared.HandleError(err)
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Getting orgs as {{.CurrentUser}}...", map[string]interface{}{
		"CurrentUser": user.Name,
	})
	cmd.UI.DisplayNewline()

	orgs, warnings, err := cmd.Actor.GetOrganizations()
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}

	if len(orgs) == 0 {
		cmd.UI.DisplayText("No orgs found.")
	} else {
		cmd.displayOrgs(orgs)
	}

	return nil
}

func (cmd OrgsCommand) displayOrgs(orgs []v2action.Organization) {
	table := [][]string{{cmd.UI.TranslateText("name")}}
	for _, org := range orgs {
		table = append(table, []string{org.Name})
	}
	cmd.UI.DisplayTableWithHeader("", table, ui.DefaultTableSpacePadding)
}
