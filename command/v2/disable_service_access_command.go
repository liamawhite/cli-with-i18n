package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DisableServiceAccessCommand struct {
	RequiredArgs    flag.Service `positional-args:"yes"`
	Organization    string       `short:"o" description:"Disable access for a specified organization"`
	ServicePlan     string       `short:"p" description:"Disable access to a specified service plan"`
	usage           interface{}  `usage:"CF_NAME disable-service-access SERVICE [-p PLAN] [-o ORG]"`
	relatedCommands interface{}  `related_commands:"marketplace, service-access, service-brokers"`
}

func (DisableServiceAccessCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DisableServiceAccessCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
