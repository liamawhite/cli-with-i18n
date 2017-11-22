package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SharePrivateDomainCommand struct {
	RequiredArgs    flag.OrgDomain `positional-args:"yes"`
	usage           interface{}    `usage:"CF_NAME share-private-domain ORG DOMAIN"`
	relatedCommands interface{}    `related_commands:"domains, unshare-private-domain"`
}

func (SharePrivateDomainCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SharePrivateDomainCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
