package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type DomainsCommand struct {
	usage           interface{} `usage:"CF_NAME domains"`
	relatedCommands interface{} `related_commands:"router-groups, create-route, routes"`
}

func (DomainsCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DomainsCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
