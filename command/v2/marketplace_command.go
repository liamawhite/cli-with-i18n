package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type MarketplaceCommand struct {
	ServicePlanInfo string      `short:"s" description:"Show plan details for a particular service offering"`
	usage           interface{} `usage:"CF_NAME marketplace [-s SERVICE]"`
	relatedCommands interface{} `related_commands:"create-service, services"`
}

func (MarketplaceCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (MarketplaceCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
