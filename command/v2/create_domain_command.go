package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type CreateDomainCommand struct {
	RequiredArgs    flag.OrgDomain `positional-args:"yes"`
	usage           interface{}    `usage:"CF_NAME create-domain ORG DOMAIN"`
	relatedCommands interface{}    `related_commands:"create-shared-domain, domains, router-groups, share-private-domain"`
}

func (CreateDomainCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (CreateDomainCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
