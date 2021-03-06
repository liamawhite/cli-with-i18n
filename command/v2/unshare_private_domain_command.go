package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type UnsharePrivateDomainCommand struct {
	RequiredArgs    flag.OrgDomain `positional-args:"yes"`
	usage           interface{}    `usage:"CF_NAME unshare-private-domain ORG DOMAIN"`
	relatedCommands interface{}    `related_commands:"delete-domain, domains"`
}

func (UnsharePrivateDomainCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (UnsharePrivateDomainCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
