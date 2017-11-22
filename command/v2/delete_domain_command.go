package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DeleteDomainCommand struct {
	RequiredArgs    flag.Domain `positional-args:"yes"`
	Force           bool        `short:"f" description:"Force deletion without confirmation"`
	usage           interface{} `usage:"CF_NAME delete-domain DOMAIN [-f]"`
	relatedCommands interface{} `related_commands:"delete-shared-domain, domains, unshare-private-domain"`
}

func (DeleteDomainCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DeleteDomainCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
